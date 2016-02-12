package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var aliasTarget = NestedResource{
	Description: "Route53 RecordSet AliasTarget",
	Properties: map[string]Schema{
		"DNSName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"EvaluateTargetHealth": Schema{
			Type: ValueBool,
		},

		"HostedZoneId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
