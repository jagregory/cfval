package schema

import (
	"fmt"

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

type IntrinsicFunctionValidate func(value parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports

func ValidateIntrinsicFunctions(value parse.IntrinsicFunction, ctx PropertyContext, supportedFunctions SupportedFunctions) (reporting.ValidateResult, reporting.Reports) {
	if !supportedFunctions[value.Key] {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "%s not valid in this location", value.Key)}
	}

	var reports reporting.Reports

	switch value.Key {
	case parse.FnAnd:
		reports = validateAnd(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnBase64:
		reports = validateBase64(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnCondition:
		reports = validateCondition(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnEquals:
		reports = validateEquals(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnFindInMap:
		reports = validateFindInMap(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnGetAtt:
		reports = validateGetAtt(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnGetAZs:
		reports = validateGetAZs(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnIf:
		reports = validateIf(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnJoin:
		reports = validateJoin(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnNot:
		reports = validateNot(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnOr:
		reports = validateOr(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnRef:
		reports = validateRef(value, PropertyContextAdd(ctx, string(value.Key)))
	case parse.FnSelect:
		reports = validateSelect(value, PropertyContextAdd(ctx, string(value.Key)))
	default:
		panic(fmt.Errorf("Unexpected Intrinsic Function %s", value.Key))
	}

	return reporting.ValidateAbort, reporting.Safe(reports)
}
