package schema

import (
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

var Period = FuncType{
	Description: "Period",

	Fn: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		num, err := strconv.Atoi(value.(string))
		if err != nil {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Period is not a number: %s", value)}
		}

		if num == 0 || num%60 != 0 {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Period is not a multiple of 60: %s", value)}
		}

		return reporting.ValidateOK, nil
	},
}
