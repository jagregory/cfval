package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
var networkInterface = NestedResource{
	Description: "EC2 NetworkInterface",
	Properties: Properties{
		"AssociatePublicIpAddress": Schema{
			Type: ValueBool,
		},

		"DeleteOnTermination": Schema{
			Type: ValueBool,
		},

		"Description": Schema{
			Type: ValueString,
		},

		"DeviceIndex": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"GroupSet": Schema{
			Array: true,
			Type:  SecurityGroupID,
		},

		"NetworkInterfaceId": Schema{
			Type: ValueString,
		},

		"PrivateIpAddress": Schema{
			Type: IPAddress,
		},

		"PrivateIpAddresses": Schema{
			Array: true,
			Type:  privateIPAddressSpecification,
		},

		"SecondaryPrivateIpAddressCount": Schema{
			Type: ValueNumber,
		},

		"SubnetId": Schema{
			Type:     SubnetID,
			Required: PropertyNotExists("NetworkInterfaceId"),
		},
	},
}
