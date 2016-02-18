package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

var JSON FuncType

func validateJSON(property Schema, value interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	switch t := value.(type) {
	case map[string]interface{}:
		return validateJSONMap(property, t, self, definitions, ctx)
	case []interface{}:
		return validateJSONArray(property, t, self, definitions, ctx)
	case string:
		return ValueString.Validate(Schema{Type: ValueString}, t, self, definitions, ctx)
	case float64:
		return ValueNumber.Validate(Schema{Type: ValueNumber}, t, self, definitions, ctx)
	case bool:
		return ValueNumber.Validate(Schema{Type: ValueBool}, t, self, definitions, ctx)
	default:
		panic(fmt.Sprintf("Unexpected JSON type %T", t))
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Value is not a JSON map", ctx.Path)}
}

func validateJSONMap(property Schema, value map[string]interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	// We pass a ValueString here as the property type so Refs etc... treat the
	// JSON as an assignable string value rather than a complex type. Bit hacky.
	builtinResult, errs := ValidateBuiltinFns(Schema{Type: ValueString}, value, self, definitions, ctx)

	if errs != nil {
		failures = append(failures, errs...)
	} else if builtinResult == reporting.ValidateAbort {
		return reporting.ValidateAbort, nil
	} else {
		for k, v := range value {
			if _, errs := validateJSON(property, v, self, definitions, ctx.Push(k)); errs != nil {
				failures = append(failures, errs...)
			}
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func validateJSONArray(property Schema, value []interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	for i, item := range value {
		if _, errs := validateJSON(property, item, self, definitions, ctx.Push(strconv.Itoa(i))); errs != nil {
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
