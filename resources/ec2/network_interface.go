package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html
func NetworkInterface() Resource {
	return Resource{
		AwsType: "AWS::EC2::NetworkInterface",

		Attributes: map[string]Schema{
			"PrimaryPrivateIpAddress": Schema{
				Type: IPAddress,
			},

			"SecondaryPrivateIpAddresses": Schema{
				Type:  IPAddress,
				Array: true,
			},
		},

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"Description": Schema{
				Type: ValueString,
			},

			"GroupSet": Schema{
				Type:  SecurityGroupID,
				Array: true,
			},

			"PrivateIpAddress": Schema{
				Type: IPAddress,
			},

			"PrivateIpAddresses": Schema{
				Type:  privateIPAddressSpecification,
				Array: true,
			},

			"SecondaryPrivateIpAddressCount": Schema{
				Type: ValueNumber,
			},

			"SourceDestCheck": Schema{
				Type: ValueBool,
			},

			"SubnetId": Schema{
				Type:     SubnetID,
				Required: constraints.Always,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}
