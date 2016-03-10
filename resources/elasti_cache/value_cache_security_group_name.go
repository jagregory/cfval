package elasti_cache

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var cacheSecurityGroupName = ConstrainedString(
	"CacheSecurityGroupName",
	func(value string, ctx PropertyContext) reporting.Reports {
		// TODO: CacheSecurityGroupName validation
		return nil
	},
)
