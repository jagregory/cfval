package schema

type Constraints []Constraint

// Constraint represents a restriction applied to property of a resource.
// e.g. between two or more properties, which cannot all be present on a
// resource at the same time.
type Constraint interface {
	// Pass tests whether a constraint has been satisfied. true means this
	// constraint has passed and validation may continue.
	Pass(map[string]interface{}) bool
}

type PropertyExists string

func (pe PropertyExists) Pass(values map[string]interface{}) bool {
	_, found := values[string(pe)]
	return found
}

type ConstraintFunc func(map[string]interface{}) bool

func (cf ConstraintFunc) Pass(values map[string]interface{}) bool {
	return cf(values)
}

func PropertyIs(prop, expected string) ConstraintFunc {
	return func(values map[string]interface{}) bool {
		if val, found := values[prop]; found {
			return val == expected
		}

		return false
	}
}

func PropertyNot(prop, notExpected string) ConstraintFunc {
	return func(values map[string]interface{}) bool {
		if val, found := values[prop]; found {
			return val != notExpected
		}

		return true
	}
}
