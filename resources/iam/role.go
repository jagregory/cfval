package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
var Role = Resource{
	AwsType: "AWS::IAM::Role",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ARN,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: map[string]Schema{
		"AssumeRolePolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"ManagedPolicyArns": Schema{
			Type: Multiple(ARN),
		},

		"Path": Schema{
			Type: ValueString,
		},

		"Policies": Schema{
			Type: Multiple(rolePolicy),
		},
	},
}
