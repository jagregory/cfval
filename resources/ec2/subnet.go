package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func Subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",
		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: AvailabilityZone,
			},

			"CidrBlock": Schema{
				Type:     CIDR,
				Required: constraints.Always,
			},

			"MapPublicIpOnLaunch": Schema{
				Type: ValueBool,
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
