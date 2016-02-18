package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

var VpcID = FuncType{
	Description: "VpcID",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, definitions, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: VpcID validation
		return reporting.ValidateOK, nil
	},
}
