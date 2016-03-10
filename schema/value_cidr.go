package schema

import (
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

const cidrPattern = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`

var CIDR = ConstrainedString(
	"CIDR",
	func(value string, ctx PropertyContext) reporting.Reports {
		if ok, _ := regexp.MatchString(cidrPattern, value); !ok {
			return reporting.Reports{reporting.NewFailure(ctx, "Cidr %s is invalid", value)}
		}

		return nil
	},
)
