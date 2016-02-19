package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources"
	"github.com/jagregory/cfval/schema"
	"github.com/mitchellh/cli"
)

type ByPath reporting.Reports

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPath) Less(i, j int) bool { return a[i].PathReadable < a[j].PathReadable }

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

func groupReports(reports reporting.Reports) map[string]reporting.Reports {
	group := make(map[string]reporting.Reports)

	for _, r := range reports {
		if items, ok := group[r.PathReadable]; ok {
			group[r.PathReadable] = append(items, r)
		} else {
			group[r.PathReadable] = reporting.Reports{r}
		}
	}

	return group
}

func printGroupedReports(reports reporting.Reports) {
	sort.Sort(ByPath(reports))
	grouped := groupReports(reports)

	for path, reports := range grouped {
		ui.Info(path)

		for _, report := range reports {
			if report.Level == reporting.Failure {
				ui.Error("  ✗ " + report.Message)
			} else if report.Level == reporting.Warning {
				ui.Warn("  ⁈ " + report.Message)
			} else {
				ui.Info("  ➜ " + report.Message)
			}
		}

		fmt.Println()
	}
}

func printReports(reports reporting.Reports) {
	sort.Sort(ByPath(reports))

	maxLength := 0
	for _, report := range reports {
		path := report.PathReadable
		if len(path) > maxLength {
			maxLength = len(path)
		}
	}

	for _, report := range reports {
		path := report.PathReadable

		str := path
		str += " "
		for i := 0; i < maxLength-len(path); i++ {
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

	if len(reports) > 0 {
		fmt.Println()
	}
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
  -format                Experimental output formatting
`
}

func (ValidateCommand) Synopsis() string {
	return "Validate a CloudFormation template"
}

func (c ValidateCommand) Run(args []string) int {
	var warningsAsErrors bool
	var forgiving bool
	var format string

	cmdFlags := flag.NewFlagSet("validate", flag.ContinueOnError)
	cmdFlags.BoolVar(&warningsAsErrors, "warnings-as-errors", false, "warnings-as-errors")
	cmdFlags.BoolVar(&forgiving, "forgiving", false, "forgiving")
	cmdFlags.StringVar(&format, "format", "oneline", "format")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
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

	template, err := parse.ParseTemplateJSON(bytes, forgiving)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 1
	}

	_, reports := schema.TemplateValidate(template, schema.NewResourceDefinitions(resources.AwsTypes))
	stats := reports.Stats()

	if format == "grouped" {
		printGroupedReports(reports)
	} else {
		printReports(reports)
	}

	if warningsAsErrors || stats.Failures > 0 {
		fmt.Printf("Fail: %d failures, %d warnings\n", stats.Failures, stats.Warnings)
		return 1
	}

	fmt.Printf("Pass: %d failures, %d warnings\n", stats.Failures, stats.Warnings)
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
