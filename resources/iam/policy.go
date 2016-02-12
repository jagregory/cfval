package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func Policy() Resource {
	return Resource{
		AwsType: "AWS::IAM::Policy",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"Groups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"PolicyDocument": Schema{
				Type:     JSON,
				Required: constraints.Always,
			},

			"PolicyName": Schema{
				Type:     ValueString,
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
}
