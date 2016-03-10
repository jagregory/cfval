package elasti_cache

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var cacheSubnetGroupName = ConstrainedString(
	"CacheSubnetGroupName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: CacheSubnetGroupName validation
		return nil
	},
)
