package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html
var RouteTable = Resource{
	AwsType: "AWS::EC2::RouteTable",

	// ID
	ReturnValue: Schema{
		Type: RouteTableID,
	},

	Properties: Properties{
		"VpcId": Schema{
			Type:     VpcID,
			Required: constraints.Always,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
