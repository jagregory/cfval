package schema

import (
	"fmt"
	"strings"
)

// Constraint represents a restriction applied to property of a resource.
// e.g. between two or more properties, which cannot all be present on a
// resource at the same time.
type Constraint interface {
	// Pass tests whether a constraint has been satisfied. true means this
	// constraint has passed and validation may continue.
	Pass(map[string]interface{}) bool

	// Describe returns a human-readable explanation of the constraint
	Describe(map[string]interface{}) string
}

type Constraints []Constraint

func (constraints Constraints) Pass(values map[string]interface{}) bool {
	for _, c := range constraints {
		if c.Pass(values) {
			return true
		}
	}

	return false
}

func (constraints Constraints) Describe(values map[string]interface{}) string {
	descriptions := make([]string, 0, len(constraints))

	for _, c := range constraints {
		if c.Pass(values) {
			descriptions = append(descriptions, c.Describe(values))
		}
	}

	return strings.Join(descriptions, " or ")
}

type PropertyExists string

func (pe PropertyExists) Pass(values map[string]interface{}) bool {
	_, found := values[string(pe)]
	return found
}

func (pe PropertyExists) Describe(values map[string]interface{}) string {
	return fmt.Sprintf("Property '%s' exists", pe)
}

type PropertyNotExists string

func (pe PropertyNotExists) Pass(values map[string]interface{}) bool {
	_, found := values[string(pe)]
	return !found
}

func (pe PropertyNotExists) Describe(values map[string]interface{}) string {
	return fmt.Sprintf("Property '%s' doens't exist", pe)
}

type ConstraintFunc struct {
	description string
	fn          func(map[string]interface{}) bool
}

func (cf ConstraintFunc) Describe(values map[string]interface{}) string {
	return cf.description
}

func (cf ConstraintFunc) Pass(values map[string]interface{}) bool {
	return cf.fn(values)
}

func PropertyIs(prop, expected string) ConstraintFunc {
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

func PropertyNot(prop, notExpected string) ConstraintFunc {
	return ConstraintFunc{
		description: fmt.Sprintf("Property '%s' shouldn't have value '%s'", prop, notExpected),
		fn: func(values map[string]interface{}) bool {
			if val, found := values[prop]; found {
				return val != notExpected
			}

			return true
		},
	}
}

type BoolConstraint bool

func (b BoolConstraint) Describe(map[string]interface{}) string {
	if bool(b) {
		return "Always"
	} else {
		return "Never"
	}
}

func (b BoolConstraint) Pass(map[string]interface{}) bool {
	return bool(b)
}

var Always Constraint = BoolConstraint(true)
var Never Constraint = BoolConstraint(false)
