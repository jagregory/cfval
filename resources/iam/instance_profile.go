package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func InstanceProfile() Resource {
	return Resource{
		AwsType: "AWS::IAM::InstanceProfile",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"Path": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Roles": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
