package reporting

import "strings"

type Reports []*Report

func (f Reports) String() string {
	failures := make([]string, len(f))
	for i := range f {
		failures[i] = f[i].String()
	}
	return strings.Join(failures, "\n")
}

func (reports Reports) Stats() Stats {
	successes := 0
	failures := 0
	warnings := 0

	for _, r := range reports {
		if r.Level == Failure {
			failures = failures + 1
		} else if r.Level == Warning {
			warnings = warnings + 1
		} else if r.Level == Success {
			successes = successes + 1
		}
	}

	return Stats{
		Failures:  failures,
		Warnings:  warnings,
		Successes: successes,
		Total:     successes + failures + warnings,
	}
}

// Safe returns either the given list of failures, or nil if there are no
// failures.
func Safe(f Reports) Reports {
	if f == nil || len(f) == 0 {
		return nil
	}

	return f
}
