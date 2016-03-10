package schema

import "github.com/jagregory/cfval/reporting"

var HostedZoneID = ConstrainedString(
	"HostedZoneID",

	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: HostedZoneID validation
		return nil
	},
)
