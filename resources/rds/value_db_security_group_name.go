package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbSecurityGroupName = FuncType{
	Description: "DBSecurityGroupName",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: DBSecurityGroupName validation
		return reporting.ValidateOK, nil
	},
}
