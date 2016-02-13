package schema

import (
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

type PropertyType interface {
	// Describe returns a human-readable description of the type in AWS, which
	// could be the AWS Resource Type or just any arbitrary description.
	Describe() string

	// Validate checks that the property is valid, including any built-in function
	// calls and stuff within the property.
	Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports)

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

type ValidateFunc func(Schema, interface{}, SelfRepresentation, []string) (reporting.ValidateResult, reporting.Reports)
type ArrayValidateFunc func([]interface{}, TemplateResource, []string) (reporting.ValidateResult, reporting.Reports)

type Schema struct {
	// Array is true when the expected value is an array of Type
	Array bool

	// Conflicts is an array of property names which cannot also be specified when
	// this property is too.
	Conflicts constraints.Constraint

	// Default is the AWS default value for this property; this is used for
	// validations when the property isn't specified
	//
	// e.g. prop X must be set to false when prop Y is true, if prop Y unspecified
	// but has a Default of true then this validation can safely fail.
	Default interface{}

	// Required is set to true if this property must have a value in the template
	Required constraints.Constraint

	// Type is the type of the Value this property is expected to contain. For
	// example "String", "Number", "JSON", or nested resources such as Tags.
	Type PropertyType

	// ValidateFunc can be used to supply a custom validation function for a
	// property for applying further validation on top of basic type checks.
	//
	// e.g. prop X must be set to false when prop Y is true
	ValidateFunc ValidateFunc

	// ArrayValidateFunc ArrayValidateFunc
}

func (s Schema) TargetType() PropertyType {
	if s.Type == nil {
		return nil
	}

	return s.Type
}

func (s Schema) Validate(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 20)

	if s.Array {
		// TODO: fix array-as-a-whole validation
		// if s.ArrayValidateFunc != nil {
		// 	if ok, errs := s.ArrayValidateFunc(value.([]interface{}), tr, context); !ok {
		// 		failures = append(failures, errs...)
		// 		pass = false
		// 	}
		// } else {
		for i, item := range value.([]interface{}) {
			if _, errs := s.validateValue(item, self, append(context, strconv.Itoa(i))); errs != nil {
				failures = append(failures, errs...)
			}
		}
		// }
	} else {
		if _, errs := s.validateValue(value, self, context); errs != nil {
			failures = append(failures, errs...)
		}
	}

	return reporting.ValidateOK, reporting.Safe(failures)
}

// validateValue takes a value and validates it against the Type of the
// current Schema and optionally runs any custom validation functions.
//
// This function is used for single value properties, and each item in array
// properties.
func (s Schema) validateValue(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)

	result, errs := s.Type.Validate(s, value, self, context)
	if result == reporting.ValidateAbort {
		// type validation instructed us to abort, so we bail with whatever failures
		// have been reported so far
		return reporting.ValidateOK, reporting.Safe(errs)
	}

	failures = append(failures, errs...)

	// run the custom validation if there is any, optionally bailing if the
	// validate tells us to, otherwise combining the results with any prior
	// failures
	if s.ValidateFunc != nil {
		result, errs := s.ValidateFunc(s, value, self, context)
		if result == reporting.ValidateAbort {
			return reporting.ValidateOK, reporting.Safe(errs)
		}

		failures = append(failures, errs...)
	}

	return reporting.ValidateOK, reporting.Safe(failures)
}
