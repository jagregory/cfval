package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
func UserToGroupAddition() Resource {
	return Resource{
		AwsType: "AWS::IAM::UserToGroupAddition",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"GroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Users": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
