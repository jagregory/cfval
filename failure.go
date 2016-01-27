package main

type Failure struct {
  Message string
  Context []string
}

func NewFailure(message string, context []string) Failure {
  return Failure{message, context}
}
