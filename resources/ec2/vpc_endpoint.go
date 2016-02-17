package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
var VPCEndpoint = Resource{
	AwsType: "AWS::EC2::VPCEndpoint",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"PolicyDocument": Schema{
			Type: JSON,
		},

		"RouteTableIds": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ServiceName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"VpcId": Schema{
			Type:     VpcID,
			Required: constraints.Always,
		},
	},
}
