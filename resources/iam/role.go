package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
func Role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",

		Attributes: map[string]Schema{
			"Arn": Schema{
				Type: ValueString,
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
				Type:  ValueString,
				Array: true,
			},

			"Path": Schema{
				Type: ValueString,
			},

			"Policies": Schema{
				Array: true,
				Type:  rolePolicy,
			},
		},
	}
}
