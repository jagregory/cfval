package constraints

import "fmt"

// PropertyExists is a constraint which will pass when a resource has a value
// for the property specified in the template.
type PropertyExists string

// Pass returns true if the property is found.
func (pe PropertyExists) Pass(values map[string]interface{}) bool {
	_, found := values[string(pe)]
	return found
}

// Describe returns a human-readable explanation of the constraint.
func (pe PropertyExists) Describe(values map[string]interface{}) string {
	return fmt.Sprintf("%s exists", pe)
}
