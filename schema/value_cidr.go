package schema

import (
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

// TODO: Can we switch this to the regexp validate?
const cidrPattern = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`

var CIDR = FuncType{
	Description: "CIDR",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		if ok, _ := regexp.MatchString(cidrPattern, value.(string)); !ok {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Cidr %s is invalid", value)}
		}

		return reporting.ValidateOK, nil
	},
}
