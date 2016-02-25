package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateBase64(base64 parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	base64Value, found := base64.UnderlyingMap["Fn::Base64"]
	if !found || base64Value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Base64\" key")}
	}

	if len(base64.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(base64.UnderlyingMap, "Fn::Base64"))}
	}

	base64ItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
	_, errs := ValueString.Validate(base64Value, base64ItemContext)
	return reporting.ValidateAbort, errs
}
