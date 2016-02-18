package schema

import "github.com/jagregory/cfval/reporting"

var VpcID = FuncType{
	Description: "VpcID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: VpcID validation
		return reporting.ValidateOK, nil
	},
}
