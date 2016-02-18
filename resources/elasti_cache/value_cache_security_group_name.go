package elasti_cache

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var cacheSecurityGroupName = FuncType{
	Description: "CacheSecurityGroupName",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: CacheSecurityGroupName validation
		return reporting.ValidateOK, nil
	},
}
