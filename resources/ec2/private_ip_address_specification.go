package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
var privateIPAddressSpecification = NestedResource{
	Description: "EC2 Network Interface Private IP Specification",
	Properties: Properties{
		"PrivateIpAddress": Schema{
			Type:     IPAddress,
			Required: Always,
		},

		"Primary": Schema{
			Type:     ValueBool,
			Required: Always,
			// TODO: You can set only one primary private IP address.
		},
	},
}
