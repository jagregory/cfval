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

func NewFailure(message string, path Path) *Report {
	return &Report{Failure, message, path.Path(), strings.Join(path.Path(), ".")}
}

func NewWarning(message string, path Path) *Report {
	return &Report{Warning, message, path.Path(), strings.Join(path.Path(), ".")}
}

func NewInvalidTypeFailure(valueType interface{}, value interface{}, path Path) *Report {
	return NewFailure(fmt.Sprintf("Property has invalid type %T, expected: %s", value, valueType), path)
}

// Safe returns either the given list of failures, or nil if there are no
// failures.
func Safe(f Reports) Reports {
	if f == nil || len(f) == 0 {
		return nil
	}

	return f
}
