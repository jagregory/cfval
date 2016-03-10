package schema

import "github.com/jagregory/cfval/reporting"

type constrainedStringValidate func(value string, ctx PropertyContext) reporting.Reports

func ConstrainedString(description string, fn constrainedStringValidate) PropertyType {
	return FuncType{
		Description: description,
		Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
			if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
				return reporting.ValidateOK, errs
			}

			return reporting.ValidateOK, fn(value.(string), ctx)
		},
	}
}

type FuncType struct {
	Description string
	Fn          func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports)
	CoercibleFn func(PropertyType) Coercion
}

func (FuncType) IsArray() bool {
	return false
}

func (ft FuncType) Describe() string {
	return ft.Description
}

func (FuncType) PropertyDefault(string) (interface{}, bool) {
	return nil, false
}

func (from FuncType) Same(to PropertyType) bool {
	if ft, ok := to.(FuncType); ok {
		return ft.Description == from.Description
	}

	return false
}

func (from FuncType) CoercibleTo(to PropertyType) Coercion {
	if from.CoercibleFn != nil {
		return from.CoercibleFn(to)
	} else if to == ValueString || to.Same(JSON) {
		return CoercionAlways
	} else if ft, ok := to.(FuncType); ok && ft.Description == from.Description {
		return CoercionAlways
	} else if to == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func (ft FuncType) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if ft.Fn == nil {
		panic("FuncType without Fn")
	}

	return ft.Fn(value, ctx)
}

func (ft FuncType) String() string {
	return ft.Description
}
