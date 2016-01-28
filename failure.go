package main

import "strings"

type Failure struct {
  Message string
  Context []string
  ContextReadable string
}

func NewFailure(message string, context []string) Failure {
  return Failure{message, context, strings.Join(context, ".")}
}
