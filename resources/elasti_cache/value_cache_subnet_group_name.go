package elasti_cache

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var cacheSubnetGroupName = FuncType{
	Description: "CacheSubnetGroupName",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: CacheSubnetGroupName validation
		return reporting.ValidateOK, nil
	},
}
