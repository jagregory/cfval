package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"flag"
)

var bail = flag.Bool("bail", false, "Bail if unrecognised resource encountered")

type Resource interface {
	Validate(t Template, context []string) (bool, []Failure)
}

func printFailures(failures []Failure) {
	maxLength := 0
	for _,failure := range failures {
		context := strings.Join(failure.Context, ".")
		if len(context) > maxLength {
			maxLength = len(context)
		}
	}

	for _,failure := range failures {
		context := strings.Join(failure.Context, ".")

		fmt.Print(context)
		fmt.Print(" ")
		for i := 0; i < maxLength - len(context); i++ {
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

	template,err := parseTemplateJSON(bytes, *bail)
	if err != nil {
    fmt.Println("Error parsing JSON:", err)
    return
  }

	if ok,errors := template.Validate(); !ok {
		printFailures(errors)
		return
	}
}
