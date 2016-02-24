package iam

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
var User = Resource{
	AwsType: "AWS::IAM::User",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ARN,
		},
	},

	// UserName
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Groups": Schema{
			Type: Multiple(ValueString),
		},

		"LoginProfile": Schema{
			Type: loginProfile,
		},

		"ManagedPolicyArns": Schema{
			Type: Multiple(ARN),
		},

		"Path": Schema{
			Type: ValueString,
		},

		"Policies": Schema{
			Type: Multiple(policy),
		},
	},
}
