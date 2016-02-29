package constraints

import "fmt"

// PropertyNot is a constraint which passes when a property of a resource
// doesn't have the expected value
func PropertyNot(prop, notExpected string) ConstraintFunc {
	return ConstraintFunc{
		description: fmt.Sprintf("Property '%s' shouldn't have value '%s'", prop, notExpected),
		fn: func(cr CurrentResource) bool {
			if val, found := cr.PropertyValueOrDefault(prop); found {
				return val != notExpected
			}

			return true
		},
	}
}
