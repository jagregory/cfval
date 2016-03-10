package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-builtin.html
func validateBase64(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnBase64, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnBase64)]
	return validateBase64Value(value, NewPropertyContext(ctx, Schema{Type: ValueString}))
}

func validateBase64Value(value interface{}, ctx PropertyContext) reporting.Reports {
	switch t := value.(type) {
	case string:
		return nil
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnIf:   true,
			parse.FnJoin: true,
			parse.FnRef:  true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid value %s", value)}
}
