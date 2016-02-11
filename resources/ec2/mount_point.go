package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
var mountPoint = NestedResource{
	Description: "EC2 MountPoint",
	Properties: Properties{
		"Device": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"VolumeId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
