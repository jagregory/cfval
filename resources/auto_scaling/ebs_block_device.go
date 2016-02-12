package auto_scaling

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
var autoScalingEbsBlockDevice = NestedResource{
	Description: "AutoScaling EBS Block Device",
	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"Encrypted": Schema{
			Type: ValueBool,
		},

		"Iops": Schema{
			Type: ValueNumber,
		},

		"SnapshotId": Schema{
			Type: ValueString,
		},

		"VolumeSize": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 1024),
		},

		"VolumeType": Schema{
			Type:    volumeType,
			Default: "standard",
		},
	},
}
