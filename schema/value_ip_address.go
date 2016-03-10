package schema

import (
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

var ipAddressRegex = regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

var IPAddress = ConstrainedString(
	"IPAddress",
	func(value string, ctx PropertyContext) reporting.Reports {
		if !ipAddressRegex.MatchString(value) {
			return reporting.Reports{
				reporting.NewFailure(ctx, "Value '%s' is not a valid IPv4 address", value),
			}
		}

		return nil
	},
)
