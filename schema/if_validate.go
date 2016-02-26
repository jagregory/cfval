package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type SupportedFunctions map[parse.IntrinsicFunctionSignature]bool

// SupportedFunctionsAll deliberately excludes Condition, because it isn't valid
// anywhere except nested in other conditions (and they cater for it explicitly)
var SupportedFunctionsAll = SupportedFunctions{
	parse.FnAnd:       true,
	parse.FnBase64:    true,
	parse.FnEquals:    true,
	parse.FnFindInMap: true,
	parse.FnGetAtt:    true,
	parse.FnGetAZs:    true,
	parse.FnIf:        true,
	parse.FnJoin:      true,
	parse.FnNot:       true,
	parse.FnOr:        true,
	parse.FnRef:       true,
	parse.FnSelect:    true,
}

func ValidateIntrinsicFunctions(value parse.IntrinsicFunction, ctx PropertyContext, supportedFunctions SupportedFunctions) (reporting.ValidateResult, reporting.Reports) {
	if !supportedFunctions[value.Key] {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "%s not valid in this location", value.Key)}
	}

	switch value.Key {
	case parse.FnAnd:
		return validateAnd(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnBase64:
		return validateBase64(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnEquals:
		return validateEquals(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnFindInMap:
		return validateFindInMap(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnGetAtt:
		return validateGetAtt(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnGetAZs:
		return validateGetAZs(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnIf:
		return validateIf(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnJoin:
		return validateJoin(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnRef:
		return validateRef(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnSelect:
		return validateSelect(value, PropertyContextAdd(ctx, string(value.Key)))
	default:
		return reporting.ValidateOK, nil
	}
}
