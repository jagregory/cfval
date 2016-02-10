package schema

import "github.com/jagregory/cfval/reporting"

var SecurityGroupName = FuncType{
	Description: "SecurityGroupName",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
		if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: SecurityGroupName validation
		return reporting.ValidateOK, nil
	},
}
