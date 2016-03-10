package schema

import "github.com/jagregory/cfval/reporting"

var KeyName = ConstrainedString(
	"KeyName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: KeyName validation
		return nil
	},
)
