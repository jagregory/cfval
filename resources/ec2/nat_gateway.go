package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html
var NatGateway = Resource{
	AwsType: "AWS::EC2::NatGateway",

	// ID
	ReturnValue: Schema{
		Type: NatGatewayID,
	},

	Properties: Properties{
		"AllocationId": Schema{
			Type:     AllocationID,
			Required: constraints.Always,
		},

		"SubnetId": Schema{
			Type: SubnetID,
      Required: constraints.Always,
		},
	},
}
