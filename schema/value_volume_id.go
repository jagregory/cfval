package schema

import "github.com/jagregory/cfval/reporting"

var VolumeID = FuncType{
	Description: "VolumeID",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: VolumeID validation
		return reporting.ValidateOK, nil
	},
}
