package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
var VPNConnectionRoute = Resource{
	AwsType: "AWS::EC2::VPNConnectionRoute",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DestinationCidrBlock": Schema{
			Type:     CIDR,
			Required: constraints.Always,
		},

		"VpnConnectionId": Schema{
			Type:     VpnConnectionID,
			Required: constraints.Always,
		},
	},
}
