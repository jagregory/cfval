package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

var HostedZoneID = FuncType{
	Description: "HostedZoneID",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, template, definitions, path); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: HostedZoneID validation
		return reporting.ValidateOK, nil
	},
}
