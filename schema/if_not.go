package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html#d0e97890
func validateNot(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnNot, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnNot)]
	switch t := value.(type) {
	case parse.IntrinsicFunction:
		itemType := Schema{Type: ValueString}
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, itemType), SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnFindInMap: true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
			parse.FnRef:       true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::Not\" key: %T", value)}
}
