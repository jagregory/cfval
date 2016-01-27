package main

import (
  "fmt"
  "strings"
)

type Failure struct {
  message, context string
}

func (f Failure) String() string {
  return fmt.Sprintf("%s\t(See: %s)", f.message, f.context)
}

func NewFailure(message string, context []string) Failure {
  return Failure{message, strings.Join(context, ".")}
}
