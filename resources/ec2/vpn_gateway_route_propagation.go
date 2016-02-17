package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
var VPNGatewayRoutePropagation = Resource{
	AwsType: "AWS::EC2::VPNGatewayRoutePropagation",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"RouteTableIds": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.Always,
		},

		"VpnGatewayId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
