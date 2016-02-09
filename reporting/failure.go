package reporting

import (
	"fmt"
	"strings"
)

type Failures []*Failure

func (f Failures) String() string {
	failures := make([]string, len(f))
	for i := range f {
		failures[i] = f[i].String()
	}
	return strings.Join(failures, "\n")
}

type Failure struct {
	Message         string
	Context         []string
	ContextReadable string
}

func (f Failure) String() string {
	return fmt.Sprintf("%s (%s)", f.Message, f.ContextReadable)
}

func NewFailure(message string, context []string) *Failure {
	return &Failure{message, context, strings.Join(context, ".")}
}

func NewInvalidTypeFailure(valueType interface{}, value interface{}, context []string) *Failure {
	return NewFailure(fmt.Sprintf("Property has invalid type %T, expected: %s", value, valueType), context)
}
