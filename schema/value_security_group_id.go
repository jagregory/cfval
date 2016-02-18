package schema

import "github.com/jagregory/cfval/reporting"

var SecurityGroupID = FuncType{
	Description: "SecurityGroupID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: SecurityGroupID validation
		return reporting.ValidateOK, nil
	},
}
