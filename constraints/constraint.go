package constraints

type CurrentResource interface {
	Properties() []string
	PropertyValue(string) (interface{}, bool)
	PropertyDefault(string) interface{}
}

// Constraint represents a restriction applied to property of a resource.
// e.g. between two or more properties, which cannot all be present on a
// resource at the same time.
type Constraint interface {
	// Pass tests whether a constraint has been satisfied. true means this
	// constraint has passed and validation may continue.
	Pass(CurrentResource) bool

	// Describe returns a human-readable explanation of the constraint.
	Describe(CurrentResource) string
}
