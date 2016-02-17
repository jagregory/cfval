package constraints

import "strings"

// All is a Constraint which only passes when all of its child constraints
// also pass.
type All []Constraint

// Pass tests whether all child constraints have been satisfied. true means all
// constraints have passed and validation may continue.
func (constraints All) Pass(values map[string]interface{}) bool {
	for _, c := range constraints {
		if !c.Pass(values) {
			return false
		}
	}

	return true
}

// Describe returns a human-readable explanation of the constraint.
func (constraints All) Describe(values map[string]interface{}) string {
	descriptions := make([]string, 0, len(constraints))

	for _, c := range constraints {
		if c.Pass(values) {
			descriptions = append(descriptions, c.Describe(values))
		}
	}

	return strings.Join(descriptions, " and ")
}
