package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func RouteTable() Resource {
	return Resource{
		AwsType: "AWS::EC2::RouteTable",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"VpcId": Schema{
				Type:     VpcID,
				Required: constraints.Always,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}
