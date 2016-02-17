package lambda

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
var Permission = Resource{
	AwsType: "AWS::Lambda::Permission",

	Properties: Properties{
		"Action": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"FunctionName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Principal": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SourceAccount": Schema{
			Type: ValueString,
		},

		"SourceArn": Schema{
			Type: ValueString,
		},
	},
}
