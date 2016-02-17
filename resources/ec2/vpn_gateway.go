package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html
var VPNGateway = Resource{
	AwsType: "AWS::EC2::VPNGateway",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Type": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: SingleValueValidate("ipsec.1"),
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},
	},
}
