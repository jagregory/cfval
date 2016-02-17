package kms

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
var Key = Resource{
	AwsType: "AWS::KMS::Key",

	// ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Description": Schema{
			Type: ValueString,
		},

		"Enabled": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"EnableKeyRotation": Schema{
			Type:    ValueBool,
			Default: false,
		},

		"KeyPolicy": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},
	},
}
