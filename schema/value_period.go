package schema

import (
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

var Period = ConstrainedString(
	"Period",
	func(value string, ctx PropertyContext) reporting.Reports {
		num, err := strconv.Atoi(value)
		if err != nil {
			return reporting.Reports{reporting.NewFailure(ctx, "Period is not a number: %s", value)}
		}

		if num == 0 || num%60 != 0 {
			return reporting.Reports{reporting.NewFailure(ctx, "Period is not a multiple of 60: %s", value)}
		}

		return nil
	},
)
