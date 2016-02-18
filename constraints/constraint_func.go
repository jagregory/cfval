package constraints

// ConstraintFunc is a shorthand for defining inline functions for constraints.
type ConstraintFunc struct {
	description string
	fn          func(CurrentResource) bool
}

// Pass just calls the func directly.
func (cf ConstraintFunc) Pass(cr CurrentResource) bool {
	return cf.fn(cr)
}

// Describe returns the description of the ConstraintFunc.
func (cf ConstraintFunc) Describe(CurrentResource) string {
	return cf.description
}
