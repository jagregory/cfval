package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-builtin.html
func validateBase64(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::Base64"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Base64\" key")}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::Base64"))}
	}

	switch t := value.(type) {
	case string:
		return reporting.ValidateAbort, nil
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnIf: true,
		})
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid value %s", value)}
}
