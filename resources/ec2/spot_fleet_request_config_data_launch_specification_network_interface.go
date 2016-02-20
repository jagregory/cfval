package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html
var spotFleetRequestConfigDataLaunchSpecificationNetworkInterface = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces",

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
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"Groups": Schema{
			Type:  SecurityGroupID,
			Array: true,
		},

		"NetworkInterfaceId": Schema{
			Type: NetworkInterfaceID,
		},

		"PrivateIpAddresses": Schema{
			Type:  spotFleetRequestConfigDataLaunchSpecificationNetworkInterfacePrivateIPAddress,
			Array: true,
		},

		"SecondaryPrivateIpAddressCount": Schema{
			Type: ValueNumber,
		},

		"SubnetId": Schema{
			Type:     SubnetID,
			Required: constraints.PropertyNotExists("NetworkInterfaceId"),
		},
	},
}
