// brain dump
// instead of having all these different ValueType's
// we should make schema.Type be an interface type which implements ??? and replaces ValidateFunc
// so instead of using EnumValidate we have a Type which knows that it can only be a certain
// combination of strings, or a Type which knows it should have an InstanceId format...

package schema

import (
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

func validateJson(value interface{}, tr TemplateResource, context []string) (reporting.ValidateResult, reporting.Failures) {
	switch t := value.(type) {
	case map[string]interface{}:
		failures := make(reporting.Failures, 0, 100)

		// TODO: Fix this up asap
		// if ok, errs := Json.validateBuiltinFns(t, tr, context); !ok && errs != nil {
		// 	failures = append(failures, errs...)
		// } else {
		// 	for key, value := range t {
		// 		if ok, errs := validateJson(value, tr, append(context, key)); !ok {
		// 			failures = append(failures, errs...)
		// 		}
		// 	}
		// }

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	case []interface{}:
		failures := make(reporting.Failures, 0, 100)

		for i, value := range t {
			if _, errs := validateJson(value, tr, append(context, strconv.Itoa(i))); errs != nil {
				failures = append(failures, errs...)
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	case string:
		return ValueString.Validate(Schema{Type: ValueString}, t, tr, context)
	case float64:
		return ValueNumber.Validate(Schema{Type: ValueNumber}, t, tr, context)
	}

	return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Value is not a JSON map", context)}
}

var Json Schema

func init() {
	Json = Schema{
		Type: ValueMap,
		// ValidateFunc: validateJson, TODO: Fixme
	}
}

type PropertyType interface {
	Describe() string
	Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures)
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
