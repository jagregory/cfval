package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type SupportedFunctions map[parse.IntrinsicFunctionSignature]bool

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
	case parse.FnBase64:
		return validateBase64(value, PropertyContextAdd(ctx, string(parse.FnBase64)))
	case parse.FnFindInMap:
		return validateFindInMap(value, PropertyContextAdd(ctx, string(parse.FnFindInMap)))
	case parse.FnGetAtt:
		return validateGetAtt(value, PropertyContextAdd(ctx, string(parse.FnGetAtt)))
	case parse.FnGetAZs:
		return validateGetAZs(value, PropertyContextAdd(ctx, string(parse.FnGetAZs)))
	case parse.FnJoin:
		return validateJoin(value, PropertyContextAdd(ctx, string(parse.FnJoin)))
	case parse.FnRef:
		return validateRef(value, PropertyContextAdd(ctx, string(parse.FnRef)))
	case parse.FnSelect:
		return validateSelect(value, PropertyContextAdd(ctx, string(parse.FnSelect)))
	default:
		return reporting.ValidateOK, nil
	}
}
