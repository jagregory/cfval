package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getavailabilityzones.html
func validateGetAZs(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnGetAZs, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnGetAZs)]
	switch t := value.(type) {
	case string:
		return nil
	case parse.IntrinsicFunction:
		azType := Schema{Type: ValueString} // TODO: Region
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, azType), SupportedFunctions{
			parse.FnRef: true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::GetAZs\" key: %s", value)}
}
