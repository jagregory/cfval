package reporting

import (
	"fmt"
	"strings"
)

type Reports []*Report

func (f Reports) String() string {
	failures := make([]string, len(f))
	for i := range f {
		failures[i] = f[i].String()
	}
	return strings.Join(failures, "\n")
}

func (reports Reports) Stats() Stats {
	failures := 0
	warnings := 0

	for _, r := range reports {
		if r.Level == Failure {
			failures = failures + 1
		} else if r.Level == Warning {
			warnings = warnings + 1
		}
	}

	return Stats{
		Failures: failures,
		Warnings: warnings,
		Total:    failures + warnings,
	}
}

type Stats struct {
	Failures, Warnings, Total int
}

type Level int

const (
	Failure Level = iota
	Warning
	Success
)

type Report struct {
	Level           Level
	Message         string
	Context         []string
	ContextReadable string
}

func (f Report) String() string {
	return fmt.Sprintf("%s (%s)", f.Message, f.ContextReadable)
}

func NewFailure(message string, context []string) *Report {
	return &Report{Failure, message, context, strings.Join(context, ".")}
}

func NewWarning(message string, context []string) *Report {
	return &Report{Warning, message, context, strings.Join(context, ".")}
}

func NewInvalidTypeFailure(valueType interface{}, value interface{}, context []string) *Report {
	return NewFailure(fmt.Sprintf("Property has invalid type %T, expected: %s", value, valueType), context)
}

// Safe returns either the given list of failures, or nil if there are no
// failures.
func Safe(f Reports) Reports {
	if f == nil || len(f) == 0 {
		return nil
	}

	return f
}
