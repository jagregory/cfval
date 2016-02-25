package schema

import (
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type ValidateFunc func(interface{}, PropertyContext) (reporting.ValidateResult, reporting.Reports)

// A Schema defines the qualities and behaviour of a property.
type Schema struct {
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
	propertyContext := NewPropertyContext(ctx, s)

	if arrayType, ok := s.Type.(ArrayPropertyType); ok {
		return validateArray(arrayType, value, propertyContext)
	} else {
		return validateValue(value, propertyContext)
	}
}

func validateArray(arrayType ArrayPropertyType, value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	itemSchema := Schema{
		Type: arrayType.Unwrap(),
	}

	switch t := value.(type) {
	case []interface{}:
		results := make(reporting.Reports, 0, 25)
		for i, item := range t {
			if _, errs := itemSchema.Validate(item, ResourceContextAdd(ctx, strconv.Itoa(i))); errs != nil {
				results = append(results, errs...)
			}
		}
		return reporting.ValidateOK, reporting.Safe(results)
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx)
	case map[string]interface{}:
		return validateMapWhereArrayShouldBe(arrayType, itemSchema, t, ctx)
	default:
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "%T used in %s property", t, arrayType.Describe())}
	}
}

// validateMapWhereArrayShouldBe runs validations against a map which was found
// where an Array was expected; this is possibly valid, and could either be a
// function reference or some attempt at coercion.
func validateMapWhereArrayShouldBe(arrayType ArrayPropertyType, itemSchema Schema, value map[string]interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if ctx.Options()[OptionExperimentMapArrayCoercion] {
		// Not sure about this behaviour, so it's behind a flag until I decide
		// whether it's wise or not.
		//
		// CloudFormation appears to allow you to flatten a single item array
		// for array properties, e.g. X: [Y] can be specified as X: Y
		//
		// So in this case if we get a map here just validate it against the
		// schema for the item of the array
		results := make(reporting.Reports, 0, 25)

		if _, errs := itemSchema.Validate(value, ctx); errs != nil {
			results = append(results, errs...)
		}

		results = append(results, reporting.NewWarning(ctx, "%s used instead of %s", arrayType.Unwrap().Describe(), arrayType.Describe()))

		return reporting.ValidateOK, reporting.Safe(results)
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "%s used instead of %s", arrayType.Unwrap().Describe(), arrayType.Describe())}
}

// validateValue takes a value and validates it against the Type of the
// current Schema and optionally runs any custom validation functions.
//
// This function is used for single value properties, and each item in array
// properties.
func validateValue(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	results := make(reporting.Reports, 0, 50)

	result, errs := ctx.Property().Type.Validate(value, ctx)
	if result == reporting.ValidateAbort {
		// type validation instructed us to abort, so we bail with whatever results
		// have been reported so far
		return reporting.ValidateOK, reporting.Safe(errs)
	}

	results = append(results, errs...)

	// run the custom validation if there is any, optionally bailing if the
	// validate tells us to, otherwise combining the results with any prior
	// results
	if fn := ctx.Property().ValidateFunc; fn != nil {
		result, errs := fn(value, ctx)
		if result == reporting.ValidateAbort {
			return reporting.ValidateOK, reporting.Safe(errs)
		}

		results = append(results, errs...)
	}

	return reporting.ValidateOK, reporting.Safe(results)
}
