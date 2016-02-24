package schema

import (
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

type ValidateFunc func(interface{}, PropertyContext) (reporting.ValidateResult, reporting.Reports)

// A Schema defines the qualities and behaviour of a property.
type Schema struct {
	// Array is true when the expected value is an array of Type
	// Array bool

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
}

func (s Schema) TargetType() PropertyType {
	if s.Type == nil {
		return nil
	}

	return s.Type
}

func (s Schema) Validate(value interface{}, ctx ResourceContext) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 20)
	propertyContext := NewPropertyContext(ctx, s)

	if arrayType, ok := s.Type.(ArrayPropertyType); ok {
		itemSchema := Schema{
			Type: arrayType.Unwrap(),
		}

		switch t := value.(type) {
		case []interface{}:
			for i, item := range t {
				if _, errs := itemSchema.Validate(item, ResourceContextAdd(ctx, strconv.Itoa(i))); errs != nil {
					failures = append(failures, errs...)
				}
			}
		case map[string]interface{}:
			// Is a map, quite possibly a Ref or something similar
			if _, errs := validateValue(t, propertyContext); errs != nil {
				failures = append(failures, errs...)
			}
		default:
			failures = append(failures, reporting.NewFailure(ctx, "%T used in %s Array property", t, s.Type.Describe()))
		}
	} else {
		if _, errs := validateValue(value, propertyContext); errs != nil {
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
func validateValue(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)

	result, errs := ctx.Property().Type.Validate(value, ctx)
	if result == reporting.ValidateAbort {
		// type validation instructed us to abort, so we bail with whatever failures
		// have been reported so far
		return reporting.ValidateOK, reporting.Safe(errs)
	}

	failures = append(failures, errs...)

	// run the custom validation if there is any, optionally bailing if the
	// validate tells us to, otherwise combining the results with any prior
	// failures
	if fn := ctx.Property().ValidateFunc; fn != nil {
		result, errs := fn(value, ctx)
		if result == reporting.ValidateAbort {
			return reporting.ValidateOK, reporting.Safe(errs)
		}

		failures = append(failures, errs...)
	}

	return reporting.ValidateOK, reporting.Safe(failures)
}
