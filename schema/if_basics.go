package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// validateIntrinsicFunctionBasicCriteria checks the essentials of an intrinsic
// function call:
//
//   1. Does it have a key matching the IF's name
//   2. Does that key have a value
//   3. No unexpected properties alongside the IF key
//
// Validation will fail if any of those criteria don't pass.
func validateIntrinsicFunctionBasicCriteria(key parse.IntrinsicFunctionSignature, builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	value, found := builtin.UnderlyingMap[string(key)]
	if !found || value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Missing \"%s\" key", key)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, string(key)))}
	}

	return nil
}
