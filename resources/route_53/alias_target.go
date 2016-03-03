package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
var aliasTarget = NestedResource{
	Description: "Route53 RecordSet AliasTarget",
	Properties: Properties{
		"DNSName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"EvaluateTargetHealth": Schema{
			Type: ValueBool,
		},

		"HostedZoneId": Schema{
			Type:     HostedZoneID,
			Required: constraints.Always,
		},
	},
}
