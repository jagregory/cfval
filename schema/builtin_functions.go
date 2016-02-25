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
	case parse.GetAtt:
		return validateGetAtt(t, PropertyContextAdd(ctx, "Fn::GetAtt"))
	case parse.Base64:
		return validateBase64(t, PropertyContextAdd(ctx, "Fn::Base64"))
	}

	// not a builtin, but this isn't necessarily bad so we don't return an error here
	return reporting.ValidateOK, nil // TODO: this really isn't clear what the intention is
}
