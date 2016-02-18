package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbSecurityGroupName = FuncType{
	Description: "DBSecurityGroupName",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, ctx Context) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: DBSecurityGroupName validation
		return reporting.ValidateOK, nil
	},
}
