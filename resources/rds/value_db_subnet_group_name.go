package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbSubnetGroupName = ConstrainedString(
	"DBSubnetGroupName",

	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: DBSubnetGroupName validation
		return nil
	},
)
