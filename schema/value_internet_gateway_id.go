package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

var InternetGatewayID = FuncType{
	Description: "InternetGatewayID",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, ctx Context) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: InternetGatewayID validation
		return reporting.ValidateOK, nil
	},
}
