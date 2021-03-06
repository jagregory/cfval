package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
var InstanceProfile = Resource{
	AwsType: "AWS::IAM::InstanceProfile",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ARN,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Path": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Roles": Schema{
			Type:     Multiple(ValueString),
			Required: constraints.Always,
		},
	},
}
