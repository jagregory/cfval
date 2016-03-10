package rds

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var dbParameterGroupName = ConstrainedString(
	"DBParameterGroupName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: DBParameterGroupName validation
		return nil
	},
)
