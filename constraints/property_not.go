package constraints

import "fmt"

type not struct {
	c Constraint
}

func Not(c Constraint) Constraint {
	return not{c}
}

func (pn not) Pass(cr CurrentResource) bool {
	return !pn.c.Pass(cr)
}

// Describe returns the description of the constraint.
func (pn not) Describe(cr CurrentResource) string {
	return pn.c.Describe(cr)
}

// PropertyNot is a constraint which passes when a property of a resource
// doesn't have the expected value
type propertyNot struct {
	prop, expected string
}

func PropertyNot(prop, expected string) propertyNot {
	return propertyNot{prop, expected}
}

// Pass will return true when the property doesn't have a value or default value
// which matches the expectation
func (pn propertyNot) Pass(cr CurrentResource) bool {
	if val, found := cr.PropertyValueOrDefault(pn.prop); found {
		return val != pn.expected
	}

	return true
}

// Describe returns the description of the constraint.
func (pn propertyNot) Describe(cr CurrentResource) string {
	value, _ := cr.PropertyValue(pn.prop)
	def, _ := cr.PropertyDefault(pn.prop)

	return fmt.Sprintf(
		"Property '%s' shouldn't have value '%s', but has {value: %v, default: %v}",
		pn.prop,
		pn.expected,
		value,
		def,
	)
}
