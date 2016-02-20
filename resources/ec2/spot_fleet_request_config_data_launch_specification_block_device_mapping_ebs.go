package ec2

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html
var spotFleetRequestConfigDataLaunchSpecificationBlockDeviceMappingEbs = NestedResource{
	Description: "SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings Ebs",

	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type: ValueBool,
		},

		"Encrypted": Schema{
			Type: ValueBool,
		},

		"Iops": Schema{
			Type: ValueNumber,
		},

		"SnapshotId": Schema{
			Type: SnapshotID,
		},

		"VolumeSize": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 1024),
		},

		"VolumeType": Schema{
			Type:    common.EbsVolumeType,
			Default: "standard",
		},
	},
}
