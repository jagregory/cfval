package schema

import "github.com/jagregory/cfval/reporting"

var SecurityGroupID = FuncType{
	Description: "SecurityGroupID",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: SecurityGroupID validation
		return reporting.ValidateOK, nil
	},
}
