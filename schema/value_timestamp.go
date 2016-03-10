package schema

import (
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

const timestampRegex = `^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$`

var Timestamp = ConstrainedString(
	"Timestamp",
	func(value string, ctx PropertyContext) reporting.Reports {
		if ok, _ := regexp.MatchString(cidrPattern, value); !ok {
			return reporting.Reports{reporting.NewFailure(ctx, "Timestamp %s is invalid", value)}
		}

		return nil
	},
)
