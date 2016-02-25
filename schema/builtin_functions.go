package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func ValidateBuiltinFns(value parse.Builtin, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch value.Key {
	case parse.BuiltinRef:
		return validateRef(value, PropertyContextAdd(ctx, string(parse.BuiltinRef)))
	case parse.BuiltinFindInMap:
		return validateFindInMap(value, PropertyContextAdd(ctx, string(parse.BuiltinFindInMap)))
	case parse.BuiltinJoin:
		return validateJoin(value, PropertyContextAdd(ctx, string(parse.BuiltinJoin)))
	case parse.BuiltinGetAtt:
		return validateGetAtt(value, PropertyContextAdd(ctx, string(parse.BuiltinGetAtt)))
	case parse.BuiltinBase64:
		return validateBase64(value, PropertyContextAdd(ctx, string(parse.BuiltinBase64)))
	default:
		return reporting.ValidateOK, nil
	}
}
