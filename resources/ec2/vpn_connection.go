package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html
var VPNConnection = Resource{
	AwsType: "AWS::EC2::VPNConnection",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Type": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"CustomerGatewayId": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"StaticRoutesOnly": Schema{
			Type: ValueBool,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},

		"VpnGatewayId": Schema{
			Type: VpnGatewayID,
		},
	},
}
