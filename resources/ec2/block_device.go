package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
var ec2EbsBlockDevice = NestedResource{
	Description: "Elastic Block Store Block Device",
	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"Encrypted": Schema{
			Type: ValueBool,
		},

		"Iops": Schema{
			Type:         ValueNumber,
			Required:     constraints.PropertyIs("VolumeType", "io1"),
			Conflicts:    constraints.PropertyNot("VolumeType", "io1"),
			ValidateFunc: IntegerRangeValidate(100, 2000),
		},

		"SnapshotId": Schema{
			Type: ValueString,
			// TODO: Required: Conditional If you specify both SnapshotId and VolumeSize, VolumeSize must be equal or greater than the size of the snapshot.
		},

		"VolumeSize": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 1024),
			// TODO: If the volume type is io1, the minimum value is 10.
			// TODO: Required: Conditional If you specify both SnapshotId and VolumeSize, VolumeSize must be equal or greater than the size of the snapshot.
		},

		"VolumeType": Schema{
			Type: EnumValue{
				Description: "VolumeType",
				Options:     []string{"io1", "gp2"},
			},
		},
	},
}
