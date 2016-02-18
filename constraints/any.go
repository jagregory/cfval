package constraints

import "strings"

// Any is a Constraint which passes when any of its child constraints pass.
type Any []Constraint

// Pass tests whether any child constraint have been satisfied. true means at
// least one constraint has passed and validation may continue.
func (constraints Any) Pass(cr CurrentResource) bool {
	for _, c := range constraints {
		if c.Pass(cr) {
			return true
		}
	}

	return false
}

// Describe returns a human-readable explanation of the constraint.
func (constraints Any) Describe(cr CurrentResource) string {
	descriptions := make([]string, 0, len(constraints))

	for _, c := range constraints {
		if c.Pass(cr) {
			descriptions = append(descriptions, c.Describe(cr))
		}
	}

	return strings.Join(descriptions, " or ")
}
