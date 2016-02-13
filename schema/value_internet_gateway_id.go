package schema

import "github.com/jagregory/cfval/reporting"

var InternetGatewayID = FuncType{
	Description: "InternetGatewayID",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: InternetGatewayID validation
		return reporting.ValidateOK, nil
	},
}
