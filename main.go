package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/jagregory/cfval/reporting"
	"github.com/mitchellh/cli"
)

type ByContext reporting.Reports

func (a ByContext) Len() int           { return len(a) }
func (a ByContext) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByContext) Less(i, j int) bool { return a[i].ContextReadable < a[j].ContextReadable }

var ui = &cli.ColoredUi{
	InfoColor:  cli.UiColorNone,
	ErrorColor: cli.UiColorRed,
	WarnColor:  cli.UiColorYellow,
	Ui: &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	},
}

func printReports(reports reporting.Reports) {
	sort.Sort(ByContext(reports))

	maxLength := 0
	for _, report := range reports {
		context := report.ContextReadable
		if len(context) > maxLength {
			maxLength = len(context)
		}
	}

	for _, report := range reports {
		context := report.ContextReadable

		str := context
		str += " "
		for i := 0; i < maxLength-len(context); i++ {
			str += "."
		}
		str += "... "
		str += report.Message

		if report.Level == reporting.Failure {
			ui.Error(str)
		} else if report.Level == reporting.Warning {
			ui.Warn(str)
		} else {
			ui.Info(str)
		}
	}
}

func printSummary(stats reporting.Stats) {
	fmt.Printf("%d failures, %d warnings\n", stats.Failures, stats.Warnings)
}

func getReadStream(args []string) (io.Reader, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return os.Stdin, nil
	} else if len(args) > 0 {
		return os.Open(args[0])
	} else {
		return nil, fmt.Errorf("Provide either a filename or pipe something")
	}
}

type ValidateCommand struct{}

func (ValidateCommand) Help() string {
	return `
Usage: cfval validate [filename]
       cat filename | cfval validate

  Given a CloudFormation JSON template, validate will parse and execute various
  tests against the template. Any problems will be printed and a non-zero exit
  code reported.

Options:

  -forgiving             Ignore unrecognised resource types
  -warnings-as-errors    Treat warnings as errors
`
}

func (ValidateCommand) Synopsis() string {
	return "Validate a CloudFormation template"
}

func (c ValidateCommand) Run(args []string) int {
	var warningsAsErrors bool
	var forgiving bool

	cmdFlags := flag.NewFlagSet("validate", flag.ContinueOnError)
	cmdFlags.BoolVar(&warningsAsErrors, "warnings-as-errors", false, "warnings-as-errors")
	cmdFlags.BoolVar(&forgiving, "forgiving", false, "forgiving")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) == 0 {
		return cli.RunResultHelp
	}

	stream, err := getReadStream(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		fmt.Println("Error reading JSON from Stdin")
		return 1
	}

	template, err := parseTemplateJSON(bytes, forgiving)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 1
	}

	fmt.Println(warningsAsErrors)

	if ok, reports := template.Validate(); !ok {
		stats := reports.Stats()

		printReports(reports)
		fmt.Println()
		printSummary(stats)

		if warningsAsErrors || stats.Failures > 0 {
			return 1
		}
	}

	return 0
}

func main() {
	app := cli.NewCLI("cfval", "0.1.0")
	app.Args = os.Args[1:]
	app.Commands = map[string]cli.CommandFactory{
		"validate": func() (cli.Command, error) {
			return ValidateCommand{}, nil
		},
	}

	exitStatus, err := app.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
