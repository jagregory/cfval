package reporting

import (
	"fmt"
	"strings"
)

type Report struct {
	Level        Level
	Message      string
	Path         []string
	PathReadable string
}

func (f Report) String() string {
	return fmt.Sprintf("%s (%s)", f.Message, f.PathReadable)
}

type Path interface {
	Path() []string
}

type Type interface {
	Describe() string
}

func NewSuccess(path Path, format string, args ...interface{}) *Report {
	return &Report{Success, fmt.Sprintf(format, args...), path.Path(), strings.Join(path.Path(), ".")}
}

func NewFailure(path Path, format string, args ...interface{}) *Report {
	return &Report{Failure, fmt.Sprintf(format, args...), path.Path(), strings.Join(path.Path(), ".")}
}

func NewWarning(path Path, format string, args ...interface{}) *Report {
	return &Report{Warning, fmt.Sprintf(format, args...), path.Path(), strings.Join(path.Path(), ".")}
}

func NewInvalidTypeFailure(path Path, valueType Type, value Type) *Report {
	return NewFailure(path, "%s used in %s property", valueType.Describe(), value.Describe())
}
