package schema

import "github.com/jagregory/cfval/reporting"

var VpcID = FuncType{
	Description: "VpcID",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: VpcID validation
		return reporting.ValidateOK, nil
	},
}
