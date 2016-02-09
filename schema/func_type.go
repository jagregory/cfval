package schema

import "github.com/jagregory/cfval/reporting"

type FuncType struct {
	Description string
	Fn          func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures)
}

func (ft FuncType) Describe() string {
	return ft.Description
}

func (ft FuncType) Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	return ft.Fn(property, value, self, context)
}

func (ft FuncType) String() string {
	return ft.Description
}
