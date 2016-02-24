package efs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
var MountTarget = Resource{
	AwsType: "AWS::EFS::MountTarget",

	// ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"FileSystemId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"IpAddress": Schema{
			Type: IPAddress,
		},

		"SecurityGroups": Schema{
			Type:     Multiple(SecurityGroupID),
			Required: constraints.Always,
		},

		"SubnetId": Schema{
			Type:     SubnetID,
			Required: constraints.Always,
		},
	},
}
