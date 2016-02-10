package schema

import "github.com/jagregory/cfval/reporting"

type FuncType struct {
	Description string
	Fn          func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures)
	CoercibleFn func(PropertyType) Coercion
}

func (ft FuncType) Describe() string {
	return ft.Description
}

func (from FuncType) CoercibleTo(to PropertyType) Coercion {
	if from.CoercibleFn != nil {
		return from.CoercibleFn(to)
	} else if to == ValueString {
		return CoercionAlways
	} else if ft, ok := to.(FuncType); ok && ft.Description == from.Description {
		return CoercionAlways
	} else if to == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func (ft FuncType) Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	return ft.Fn(property, value, self, context)
}

func (ft FuncType) String() string {
	return ft.Description
}
