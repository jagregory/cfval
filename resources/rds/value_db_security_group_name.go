package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbSecurityGroupName = ConstrainedString(
	"DBSecurityGroupName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: DBSecurityGroupName validation
		return nil
	},
)
