package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html#d0e97890
func validateNot(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::Not"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Not\" key")}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::Not"))}
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnFindInMap: true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::Not\" key: %T", value)}
}
