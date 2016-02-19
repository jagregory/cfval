package lambda

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
var Function = Resource{
	AwsType: "AWS::Lambda::Function",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ARN,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Code": Schema{
			Type:     code,
			Required: constraints.Always,
		},

		"Description": Schema{
			Type: ValueString,
		},

		"Handler": Schema{
			Type:     ValueString,
			Required: constraints.Always,
			ValidateFunc: RegexpValidate(
				`^[^\s]{0,128}$`,
				"Minimum length of 0. Maximum length of 128.",
			),
		},

		"MemorySize": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(128, 1536),
		},

		"Role": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"Runtime": Schema{
			Type: EnumValue{
				Description: "Runtime",
				Options:     []string{"nodejs", "java8", "python2.7"},
			},
			Required: constraints.Always,
		},

		"Timeout": Schema{
			Type:    ValueNumber,
			Default: 3,
		},
	},
}
