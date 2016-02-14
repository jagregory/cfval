package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html
var spotFleetRequestConfigDataLaunchSpecificationNetworkInterfacePrivateIPAddress = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces PrivateIpAddresses",

	Properties: Properties{
		"Primary": Schema{
			Type: ValueBool,
		},

		"PrivateIpAddress": Schema{
			Type:     IPAddress,
			Required: constraints.Always,
		},
	},
}
