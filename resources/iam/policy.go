package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
var Policy = Resource{
	AwsType: "AWS::IAM::Policy",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Groups": Schema{
			Type: Multiple(ValueString),
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
			Type: Multiple(ValueString),
			Required: constraints.All{
				constraints.PropertyNotExists("Groups"),
				constraints.PropertyNotExists("Users"),
			},
		},

		"Users": Schema{
			Type: Multiple(ValueString),
			Required: constraints.All{
				constraints.PropertyNotExists("Groups"),
				constraints.PropertyNotExists("Roles"),
			},
		},
	},
}
