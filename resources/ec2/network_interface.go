package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html
var NetworkInterface = Resource{
	AwsType: "AWS::EC2::NetworkInterface",

	Attributes: map[string]Schema{
		"PrimaryPrivateIpAddress": Schema{
			Type: IPAddress,
		},

		"SecondaryPrivateIpAddresses": Schema{
			Type: Multiple(IPAddress),
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
			Type: Multiple(SecurityGroupID),
		},

		"PrivateIpAddress": Schema{
			Type: IPAddress,
		},

		"PrivateIpAddresses": Schema{
			Type: Multiple(privateIPAddressSpecification),
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
			Type: Multiple(common.ResourceTag),
		},
	},
}
