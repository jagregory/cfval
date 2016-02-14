package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html
var hostedZoneVPC = NestedResource{
	Description: "Route 53 HostedZoneVPC",

	Properties: Properties{
		"VPCId": Schema{
			Type:     VpcID,
			Required: constraints.Always,
		},

		"VPCRegion": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
