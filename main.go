package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
  bytes, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    fmt.Println("Error reading JSON from Stdin")
    return
  }

	template,err := parseTemplateJSON(bytes)
	if err != nil {
    fmt.Println("Error parsing JSON")
    return
  }

	if ok,errors := template.Validate(); !ok {
		printFailures(errors)
		return
	}
}
