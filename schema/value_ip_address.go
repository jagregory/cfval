package schema

import "github.com/jagregory/cfval/reporting"

var IPAddress = FuncType{
	Description: "IPAddress",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
		if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: IPAddress validation
		return reporting.ValidateOK, nil
	},
}
