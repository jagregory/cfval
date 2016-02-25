package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func ValidateBuiltinFns(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch t := value.(type) {
	case parse.Ref:
		return validateRef(t, PropertyContextAdd(ctx, "Ref"))
	case parse.FindInMap:
		return validateFindInMap(t, PropertyContextAdd(ctx, "Fn::FindInMap"))
	case map[string]interface{}:
		if join, ok := t["Fn::Join"]; ok {
			return validateJoin(join, PropertyContextAdd(ctx, "Fn::Join"))
		}

		if getatt, ok := t["Fn::GetAtt"]; ok {
			if arr, ok := getatt.([]interface{}); ok {
				return NewGetAtt(arr).Validate(PropertyContextAdd(ctx, "GetAtt"))
			}

			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt must be an array")}
		}

		if base64, ok := t["Fn::Base64"]; ok {
			return validateBase64(base64, PropertyContextAdd(ctx, "Fn::Base64"))
		}
	}

	// not a builtin, but this isn't necessarily bad so we don't return an error here
	return reporting.ValidateOK, nil // TODO: this really isn't clear what the intention is
}

func validateBase64(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	base64ItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
	_, errs := ValueString.Validate(value, base64ItemContext)
	return reporting.ValidateAbort, errs
}

func validateJoin(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if items, ok := value.([]interface{}); ok {
		if len(items) != 2 {
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join has incorrect number of arguments (expected: 2, actual: %d)", len(items))}
		}

		_, ok := items[0].(string)
		if !ok {
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join '%s' is not a valid delimiter", items[0])}
		}

		parts, ok := items[1].([]interface{})
		if !ok {
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join items are not valid: %s", items[1])}
		}

		joinItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
		failures := make(reporting.Reports, 0, len(parts))
		for i, part := range parts {
			if _, errs := ValueString.Validate(part, PropertyContextAdd(joinItemContext, "1", strconv.Itoa(i))); errs != nil {
				failures = append(failures, errs...)
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateAbort, nil
		}

		return reporting.ValidateAbort, failures
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt has invalid value '%s'", value)}
}
