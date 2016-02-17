package constraints

// ConstraintFunc is a shorthand for defining inline functions for constraints.
type ConstraintFunc struct {
	description string
	fn          func(map[string]interface{}) bool
}

// Pass just calls the func directly.
func (cf ConstraintFunc) Pass(values map[string]interface{}) bool {
	return cf.fn(values)
}

// Describe returns the description of the ConstraintFunc.
func (cf ConstraintFunc) Describe(values map[string]interface{}) string {
	return cf.description
}
