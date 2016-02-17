package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
var ManagedPolicy = Resource{
	AwsType: "AWS::IAM::ManagedPolicy",

	// Arn
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Description": Schema{
			Type: ValueString,
		},

		"Groups": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Path": Schema{
			Type:    ValueString,
			Default: "/",
		},

		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"Roles": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Users": Schema{
			Type:  ValueString,
			Array: true,
		},
	},
}
