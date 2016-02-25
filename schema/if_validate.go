package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func ValidateIntrinsicFunctions(value parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
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
	default:
		return reporting.ValidateOK, nil
	}
}
