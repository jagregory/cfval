package schema

import "github.com/jagregory/cfval/reporting"

var KeyName = FuncType{
	Description: "KeyName",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: KeyName validation
		return reporting.ValidateOK, nil
	},
}
