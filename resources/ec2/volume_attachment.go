package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
var VolumeAttachment = Resource{
	AwsType: "AWS::EC2::VolumeAttachment",

	Properties: Properties{
		"Device": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"InstanceId": Schema{
			Type:     InstanceID,
			Required: constraints.Always,
		},

		"VolumeId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
