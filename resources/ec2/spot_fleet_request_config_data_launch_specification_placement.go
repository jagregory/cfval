package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html
var spotFleetRequestConfigDataLaunchSpecificationPlacement = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications Placement",

	Properties: Properties{
		"AvailabilityZone": Schema{
			Type: AvailabilityZone,
		},

		"GroupName": Schema{
			Type: ValueString,
		},
	},
}
