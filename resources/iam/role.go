package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func Role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",

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
				Type: NestedResource{
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
				},
			},
		},
	}
}
