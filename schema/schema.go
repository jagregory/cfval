package schema

import (
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

type PropertyType interface {
	Describe() string
	Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures)

	// CoercibleTo will return true for types which the value of this property can
	// be coerced into. e.g. A number can be coerced to a string
	// CoercionAlways means a type is always coercible to another
	// 	 e.g. all numbers are valid strings
	// CoercionNever means a type is never coercible to another
	//   e.g. a number is never a valid bool
	// CoercionBegrudgingly means a type can be coerced but results may vary
	//   e.g. a string can be coerced to a number, but only if it is numerically
	//        valid.
	//
	// CoerceAlways and CoercionBegrudgingly are equivalent right now, but in
	// future a warning may be issued for begrudging conversions.
	CoercibleTo(PropertyType) Coercion
}

type ValidateFunc func(interface{}, TemplateResource, []string) (reporting.ValidateResult, reporting.Failures)
type ArrayValidateFunc func([]interface{}, TemplateResource, []string) (reporting.ValidateResult, reporting.Failures)

type Schema struct {
	Array          bool
	Conflicts      []string
	Default        interface{}
	Required       bool
	RequiredIf     []string
	RequiredUnless []string
	Type           PropertyType
	// ValidateFunc      ValidateFunc
	// ArrayValidateFunc ArrayValidateFunc
}

func (s Schema) TargetType() ValueType {
	if t, ok := s.Type.(ValueType); ok {
		return t
	}

	return ValueUnknown
}

func (s Schema) Validate(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if !s.Required && value == nil {
		return reporting.ValidateOK, nil
	} else if s.Required && value == nil {
		return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Required property is missing", context)}
	}

	failures := make(reporting.Failures, 0, 20)

	if s.Array {
		// TODO: fixme
		// if s.ArrayValidateFunc != nil {
		// 	if ok, errs := s.ArrayValidateFunc(value.([]interface{}), tr, context); !ok {
		// 		failures = append(failures, errs...)
		// 		pass = false
		// 	}
		// } else {
		for i, item := range value.([]interface{}) {
			if _, errs := s.Type.Validate(s, item, self, append(context, strconv.Itoa(i))); errs != nil {
				failures = append(failures, errs...)
			}
		}
		// }
	} else {
		if _, errs := s.Type.Validate(s, value, self, context); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}
