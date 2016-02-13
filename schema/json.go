package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

var JSON FuncType

func validateJSON(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
	switch t := value.(type) {
	case map[string]interface{}:
		return validateJSONMap(property, t, self, context)
	case []interface{}:
		return validateJSONArray(property, t, self, context)
	case string:
		return ValueString.Validate(Schema{Type: ValueString}, t, self, context)
	case float64:
		return ValueNumber.Validate(Schema{Type: ValueNumber}, t, self, context)
	case bool:
		return ValueNumber.Validate(Schema{Type: ValueBool}, t, self, context)
	default:
		panic(fmt.Sprintf("Unexpected JSON type %T", t))
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Value is not a JSON map", context)}
}

func validateJSONMap(property Schema, value map[string]interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	// We pass a ValueString here as the property type so Refs etc... treat the
	// JSON as an assignable string value rather than a complex type. Bit hacky.
	builtinResult, errs := ValidateBuiltinFns(Schema{Type: ValueString}, value, self, context)

	if errs != nil {
		failures = append(failures, errs...)
	} else if builtinResult == reporting.ValidateAbort {
		return reporting.ValidateAbort, nil
	} else {
		for k, v := range value {
			if _, errs := validateJSON(property, v, self, append(context, k)); errs != nil {
				failures = append(failures, errs...)
			}
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func validateJSONArray(property Schema, value []interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	for i, item := range value {
		if _, errs := validateJSON(property, item, self, append(context, strconv.Itoa(i))); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func init() {
	JSON = FuncType{
		Description: "JSON",

		Fn: validateJSON,
	}
}
