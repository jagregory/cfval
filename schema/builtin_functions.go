package schema

import (
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

func ValidateBuiltinFns(value map[string]interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if ref, ok := value["Ref"]; ok {
		if str, ok := ref.(string); ok {
			return NewRef(str).Validate(PropertyContextAdd(ctx, "Ref"))
		}

		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Ref must be a string")}
	}

	if join, ok := value["Fn::Join"]; ok {
		return validateJoin(join, PropertyContextAdd(ctx, "Fn::Join"))
	}

	if getatt, ok := value["Fn::GetAtt"]; ok {
		if arr, ok := getatt.([]interface{}); ok {
			return NewGetAtt(arr).Validate(PropertyContextAdd(ctx, "GetAtt"))
		}

		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt must be an array")}
	}

	if find, ok := value["Fn::FindInMap"]; ok {
		return validateFindInMap(find, PropertyContextAdd(ctx, "Fn::FindInMap"))
	}

	if base64, ok := value["Fn::Base64"]; ok {
		return validateBase64(base64, PropertyContextAdd(ctx, "Fn::Base64"))
	}

	// not a builtin, but this isn't necessarily bad so we don't return an error here
	return reporting.ValidateOK, nil
}

// TODO: Supported functions within a function
func validateFindInMap(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	find, ok := value.([]interface{})
	if !ok {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Options need to be an array")}
	}

	if len(find) != 3 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Options has wrong number of items, expected: 3, actual: %d", len(find))}
	}

	findInMapItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})

	mapName := find[0]
	_, mapNameIsString := mapName.(string)
	if _, errs := ValueString.Validate(mapName, PropertyContextAdd(findInMapItemContext, "0")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString {
		// map name is a string, so we can do some further interrogation
		// TODO: lookup whether MapName is a valid Map
	}

	topLevelKey := find[1]
	_, topLevelKeyIsString := topLevelKey.(string)
	if _, errs := ValueString.Validate(topLevelKey, PropertyContextAdd(findInMapItemContext, "1")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString {
		// TODO: lookup whether topLevelKey is in mapName
	}

	secondLevelKey := find[2]
	_, secondLevelKeyIsString := secondLevelKey.(string)
	if _, errs := ValueString.Validate(secondLevelKey, PropertyContextAdd(findInMapItemContext, "2")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString && secondLevelKeyIsString {
		// TODO: lookup whether secondLevelKeyIsString is in topLevelKey
	}

	return reporting.ValidateAbort, nil
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
