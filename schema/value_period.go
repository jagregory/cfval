package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

var Period = FuncType{
	Description: "Period",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, definitions, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		num, err := strconv.Atoi(value.(string))
		if err != nil {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Period is not a number: %s", value), ctx.Path)}
		}

		if num == 0 || num%60 != 0 {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Period is not a multiple of 60: %s", value), ctx.Path)}
		}

		return reporting.ValidateOK, nil
	},
}
