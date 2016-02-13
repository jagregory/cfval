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

var forgiving = flag.Bool("forgiving", false, "Ignore unrecognised resources")

type ByContext reporting.Failures

func (a ByContext) Len() int           { return len(a) }
func (a ByContext) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByContext) Less(i, j int) bool { return a[i].ContextReadable < a[j].ContextReadable }

func printFailures(failures reporting.Failures) {
	sort.Sort(ByContext(failures))

	maxLength := 0
	for _, failure := range failures {
		context := failure.ContextReadable
		if len(context) > maxLength {
			maxLength = len(context)
		}
	}

	for _, failure := range failures {
		context := failure.ContextReadable

		fmt.Print(context)
		fmt.Print(" ")
		for i := 0; i < maxLength-len(context); i++ {
			fmt.Print(".")
		}
		fmt.Print("... ")
		fmt.Printf("%s\n", failure.Message)
	}
}

func printSummary(failures reporting.Failures) {
	fmt.Printf("%d failures\n", len(failures))
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
`
}

func (ValidateCommand) Synopsis() string {
	return "Validate a CloudFormation template"
}

func (ValidateCommand) Run(args []string) int {
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

	template, err := parseTemplateJSON(bytes, *forgiving)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 1
	}

	if ok, errors := template.Validate(); !ok {
		printFailures(errors)
		fmt.Println()
		printSummary(errors)
		return 1
	}

	fmt.Println("Pass, no errors found.")
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
