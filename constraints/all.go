package constraints

import "strings"

// All is a Constraint which only passes when all of its child constraints
// also pass.
type All []Constraint

// Pass tests whether all child constraints have been satisfied. true means all
// constraints have passed and validation may continue.
func (constraints All) Pass(cr CurrentResource) bool {
	for _, c := range constraints {
		if !c.Pass(cr) {
			return false
		}
	}

	return true
}

// Describe returns a human-readable explanation of the constraint.
func (constraints All) Describe(cr CurrentResource) string {
	descriptions := make([]string, 0, len(constraints))

	for _, c := range constraints {
		if c.Pass(cr) {
			descriptions = append(descriptions, c.Describe(cr))
		}
	}

	return strings.Join(descriptions, " and ")
}
