package reporting

import (
	"fmt"
	"strings"
)

type Failures []*Failure

type Failure struct {
	Message         string
	Context         []string
	ContextReadable string
}

func NewFailure(message string, context []string) *Failure {
	return &Failure{message, context, strings.Join(context, ".")}
}

func NewInvalidTypeFailure(valueType interface{}, value interface{}, context []string) *Failure {
	return NewFailure(fmt.Sprintf("Property has invalid type %T, expected: %s", value, valueType), context)
}
