package constraints

import "fmt"

// PropertyNotExists is a constraint which will pass when a resource doesn't
// have a value specified in the template.
type PropertyNotExists string

// Pass will return true when the property doesn't exist on the resource.
func (pe PropertyNotExists) Pass(cr CurrentResource) bool {
	_, found := cr.PropertyValue(string(pe))
	return !found
}

// Describe returns a human-readable explanation of the constraint.
func (pe PropertyNotExists) Describe(CurrentResource) string {
	return fmt.Sprintf("%s doesn't exist", pe)
}
