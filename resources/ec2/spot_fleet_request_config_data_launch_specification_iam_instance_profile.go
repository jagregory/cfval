package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html
var spotFleetRequestConfigDataLaunchSpecificationIamInstanceProfile = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications IamInstanceProfile",

	Properties: Properties{
		"Arn": Schema{
			Type: ARN,
		},
	},
}
