package schema

import "github.com/jagregory/cfval/reporting"

var SecurityGroupName = ConstrainedString(
	"SecurityGroupName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: SecurityGroupName validation
		return nil
	},
)
