package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var forgiving = flag.Bool("forgiving", false, "Ignore unrecognised resources")

type ByContext []Failure

func (a ByContext) Len() int           { return len(a) }
func (a ByContext) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByContext) Less(i, j int) bool { return a[i].ContextReadable < a[j].ContextReadable }

func printFailures(failures []Failure) {
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

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading JSON from Stdin")
		return
	}

	template, err := parseTemplateJSON(bytes, *forgiving)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if ok, errors := template.Validate(); !ok {
		printFailures(errors)
		return
	}
}
