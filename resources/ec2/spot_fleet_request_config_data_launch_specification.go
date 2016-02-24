package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html
var spotFleetRequestConfigDataLaunchSpecification = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications",

	Properties: Properties{
		"BlockDeviceMappings": Schema{
			Type: Multiple(spotFleetRequestConfigDataLaunchSpecificationBlockDeviceMapping),
		},

		"EbsOptimized": Schema{
			Type: ValueBool,
		},

		"IamInstanceProfile": Schema{
			Type: spotFleetRequestConfigDataLaunchSpecificationIamInstanceProfile,
		},

		"ImageId": Schema{
			Type:     ImageID,
			Required: constraints.Always,
		},

		"InstanceType": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"KernelId": Schema{
			Type: ValueString,
		},

		"KeyName": Schema{
			Type: KeyName,
		},

		"Monitoring": Schema{
			Type: spotFleetRequestConfigDataLaunchSpecificationMonitoring,
		},

		"NetworkInterfaces": Schema{
			Type: Multiple(spotFleetRequestConfigDataLaunchSpecificationNetworkInterface),
		},

		"Placement": Schema{
			Type: spotFleetRequestConfigDataLaunchSpecificationPlacement,
		},

		"RamdiskId": Schema{
			Type: ValueString,
		},

		"SecurityGroups": Schema{
			Type: Multiple(spotFleetRequestConfigDataLaunchSpecificationSecurityGroup),
		},

		"SubnetId": Schema{
			Type: SubnetID,
		},

		"UserData": Schema{
			Type: ValueString,
		},

		"WeightedCapacity": Schema{
			Type: ValueNumber,
		},
	},
}
