package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
var NetworkACL = Resource{
	AwsType: "AWS::EC2::NetworkAcl",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
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
