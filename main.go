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

type ByGrouping reporting.Reports

func (a ByGrouping) Len() int      { return len(a) }
func (a ByGrouping) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByGrouping) Less(i, j int) bool {
	if a[i].Level < a[j].Level {
		return true
	} else if a[i].Message < a[j].Message {
		return true
	}

	return false
}

var ui = &cli.ColoredUi{
	InfoColor:   cli.UiColorGreen,
	ErrorColor:  cli.UiColorRed,
	WarnColor:   cli.UiColorYellow,
	OutputColor: cli.UiColor{Code: cli.UiColorNone.Code, Bold: true},
	Ui: &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	},
}

func groupReports(reports reporting.Reports) (map[string]reporting.Reports, []string) {
	group := make(map[string]reporting.Reports)

	for _, r := range reports {
		if items, ok := group[r.PathReadable]; ok {
			group[r.PathReadable] = append(items, r)
		} else {
			group[r.PathReadable] = reporting.Reports{r}
		}
	}

	paths := make([]string, 0, len(group))
	for key, _ := range group {
		paths = append(paths, key)
	}
	sort.StringSlice(paths).Sort()

	return group, paths
}

func filterReportsByLevel(reports reporting.Reports, level string) reporting.Reports {
	if level == "all" {
		return reports
	}

	filtered := make(reporting.Reports, 0, len(reports))

	for _, r := range reports {
		if (level == "warning" && (r.Level == reporting.Warning || r.Level == reporting.Failure)) || (level == "failure" && r.Level == reporting.Failure) {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

func printGroupedReports(reports reporting.Reports) {
	groupedItems, order := groupReports(reports)

	for _, path := range order {
		reports := groupedItems[path]
		sort.Sort(ByGrouping(reports))
		ui.Output(path)

		for _, report := range reports {
			if report.Level == reporting.Failure {
				ui.Error("  ✗ " + report.Message)
			} else if report.Level == reporting.Warning {
				ui.Warn("  ⁈ " + report.Message)
			} else {
				ui.Info("  ✓ " + report.Message)
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
  -format                Output formatting [experimental]
  -level                 Output level [experimental]
  -warnings-as-errors    Treat warnings as errors
`
}

func (ValidateCommand) Synopsis() string {
	return "Validate a CloudFormation template"
}

func (c ValidateCommand) Run(args []string) int {
	var warningsAsErrors bool
	var forgiving bool
	var format string
	var level string

	cmdFlags := flag.NewFlagSet("validate", flag.ContinueOnError)
	cmdFlags.BoolVar(&warningsAsErrors, "warnings-as-errors", false, "warnings-as-errors")
	cmdFlags.BoolVar(&forgiving, "forgiving", false, "forgiving")
	cmdFlags.StringVar(&format, "format", "oneline", "format")
	cmdFlags.StringVar(&level, "level", "warning", "level")
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

	template, err := parse.ParseTemplateJSON(bytes, forgiving)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 1
	}

	_, reports := schema.TemplateValidate(template, schema.NewResourceDefinitions(resources.AwsTypes))
	stats := reports.Stats()

	if format == "grouped" {
		printGroupedReports(filterReportsByLevel(reports, level))
	} else {
		printReports(filterReportsByLevel(reports, level))
	}

	if warningsAsErrors || stats.Failures > 0 {
		fmt.Printf("Fail: %d pass, %d fail, %d warn\n", stats.Successes, stats.Failures, stats.Warnings)
		return 1
	}

	fmt.Printf("Pass: %d pass, %d fail, %d warn\n", stats.Successes, stats.Failures, stats.Warnings)
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
