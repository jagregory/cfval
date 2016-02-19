package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html
var Volume = Resource{
	AwsType: "AWS::EC2::Volume",

	// ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AutoEnableIO": Schema{
			Type: ValueBool,
		},

		"AvailabilityZone": Schema{
			Type:     AvailabilityZone,
			Required: constraints.Always,
		},

		"Encrypted": Schema{
			Type:     ValueBool,
			Required: constraints.PropertyExists("KmsKeyId"),
		},

		"Iops": Schema{
			Type:         ValueNumber,
			Required:     constraints.PropertyIs("VolumeType", "io1"),
			ValidateFunc: IntegerRangeValidate(1, 4000),
		},

		"KmsKeyId": Schema{
			Type: ARN,
		},

		"Size": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("SnapshotId"),
		},

		"SnapshotId": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},

		"VolumeType": Schema{
			Type:     common.EbsVolumeType,
			Array:    true,
			Required: constraints.Always,
		},
	},
}
