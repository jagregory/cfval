package constraints

import "fmt"

// PropertyIs is a constraint which passes when a property exists on a resource
// and has the expected value or default value.
func PropertyIs(prop string, expected interface{}) ConstraintFunc {
	return ConstraintFunc{
		description: fmt.Sprintf("Property '%s' has value '%s'", prop, expected),
		fn: func(values map[string]interface{}) bool {
			if val, found := values[prop]; found {
				return val == expected
			}

			return false
		},
	}
}
