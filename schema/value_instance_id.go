package schema

import "github.com/jagregory/cfval/reporting"

var InstanceID = FuncType{
	Description: "InstanceID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: InstanceID validation
		return reporting.ValidateOK, nil
	},
}
