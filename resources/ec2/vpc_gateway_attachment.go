package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html
var VPCGatewayAttachment = Resource{
	AwsType: "AWS::EC2::VPCGatewayAttachment",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"InternetGatewayId": Schema{
			Type:      InternetGatewayID,
			Required:  constraints.PropertyNotExists("VpnGatewayId"),
			Conflicts: constraints.PropertyExists("VpnGatewayId"),
		},

		"VpcId": Schema{
			Required: constraints.Always,
			Type:     VpcID,
		},

		"VpnGatewayId": Schema{
			Type:      ValueString,
			Required:  constraints.PropertyNotExists("InternetGatewayId"),
			Conflicts: constraints.PropertyExists("InternetGatewayId"),
		},
	},
}
