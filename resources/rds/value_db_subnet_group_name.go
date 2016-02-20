package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbSubnetGroupName = FuncType{
	Description: "DBSubnetGroupName",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: DBSubnetGroupName validation
		return reporting.ValidateOK, nil
	},
}
