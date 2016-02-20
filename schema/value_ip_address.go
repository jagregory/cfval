package schema

import (
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

var ipAddressRegex = regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

var IPAddress = FuncType{
	Description: "IPAddress",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		if !ipAddressRegex.MatchString(value.(string)) {
			return reporting.ValidateOK, reporting.Reports{
				reporting.NewFailure(ctx, "Value '%s' is not a valid IPv4 address", value),
			}
		}

		return reporting.ValidateOK, nil
	},
}
