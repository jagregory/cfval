package schema

import "github.com/jagregory/cfval/reporting"

var SubnetID = FuncType{
	Description: "SubnetID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: SubnetID validation
		return reporting.ValidateOK, nil
	},
}
