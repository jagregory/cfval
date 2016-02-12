package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
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
				Required: constraints.All{
					constraints.PropertyNotExists("Roles"),
					constraints.PropertyNotExists("Users"),
				},
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
				Required: constraints.All{
					constraints.PropertyNotExists("Groups"),
					constraints.PropertyNotExists("Users"),
				},
			},

			"Users": Schema{
				Type:  ValueString,
				Array: true,
				Required: constraints.All{
					constraints.PropertyNotExists("Groups"),
					constraints.PropertyNotExists("Roles"),
				},
			},
		},
	}
}
