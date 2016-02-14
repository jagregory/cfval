package iam

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
func User() Resource {
	return Resource{
		AwsType: "AWS::IAM::User",

		Attributes: map[string]Schema{
			"Arn": Schema{
				Type: ValueString,
			},
		},

		// UserName
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"Groups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"LoginProfile": Schema{
				Type: loginProfile,
			},

			"ManagedPolicyArns": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Path": Schema{
				Type: ValueString,
			},

			"Policies": Schema{
				Type:  policy,
				Array: true,
			},
		},
	}
}
