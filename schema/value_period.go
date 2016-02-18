package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

var Period = FuncType{
	Description: "Period",

	Fn: func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, template, definitions, path); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		num, err := strconv.Atoi(value.(string))
		if err != nil {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Period is not a number: %s", value), path)}
		}

		if num == 0 || num%60 != 0 {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Period is not a multiple of 60: %s", value), path)}
		}

		return reporting.ValidateOK, nil
	},
}
