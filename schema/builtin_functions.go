package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func ValidateBuiltinFns(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch t := value.(type) {
	case parse.Ref:
		return validateRef(t, PropertyContextAdd(ctx, "Ref"))
	case parse.FindInMap:
		return validateFindInMap(t, PropertyContextAdd(ctx, "Fn::FindInMap"))
	case parse.Join:
		return validateJoin(t, PropertyContextAdd(ctx, "Fn::Join"))
	case map[string]interface{}:
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
