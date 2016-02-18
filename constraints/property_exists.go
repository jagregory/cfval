package constraints

import "fmt"

// PropertyExists is a constraint which will pass when a resource has a value
// for the property specified in the template.
type PropertyExists string

// Pass returns true if the property is found.
func (pe PropertyExists) Pass(cr CurrentResource) bool {
	_, found := cr.PropertyValue(string(pe))
	return found
}

// Describe returns a human-readable explanation of the constraint.
func (pe PropertyExists) Describe(CurrentResource) string {
	return fmt.Sprintf("%s exists", pe)
}
