package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html
var spotFleetRequestConfigDataLaunchSpecificationSecurityGroup = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications SecurityGroups",

	Properties: Properties{
		"GroupId": Schema{
			Type: SecurityGroupID,
		},
	},
}
