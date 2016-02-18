package schema

import "github.com/jagregory/cfval/reporting"

var InternetGatewayID = FuncType{
	Description: "InternetGatewayID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: InternetGatewayID validation
		return reporting.ValidateOK, nil
	},
}
