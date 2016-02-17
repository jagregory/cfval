package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbParameterGroupName = FuncType{
	Description: "DBParameterGroupName",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, definitions, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: DBParameterGroupName validation
		return reporting.ValidateOK, nil
	},
}
