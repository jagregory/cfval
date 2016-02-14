package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
func VPCPeeringConnection() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCPeeringConnection",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"PeerVpcId": Schema{
				Type:     VpcID,
				Required: constraints.Always,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VpcId": Schema{
				Type:     VpcID,
				Required: constraints.Always,
			},
		},
	}
}
