package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html
var spotFleetRequestConfigDataLaunchSpecificationMonitoring = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications Monitoring",

	Properties: Properties{
		"Enabled": Schema{
			Type: ValueBool,
		},
	},
}
