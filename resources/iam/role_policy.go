package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
var rolePolicy = NestedResource{
	Description: "IAM Role Policy",
	Properties: map[string]Schema{
		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"PolicyName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
