package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
var loginProfile = NestedResource{
	Description: "IAM User LoginProfile",

	Properties: Properties{
		"Password": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PasswordResetRequired": Schema{
			Type: ValueBool,
		},
	},
}
