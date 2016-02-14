package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
func AccessKey() Resource {
	return Resource{
		AwsType: "AWS::IAM::AccessKey",

		Attributes: map[string]Schema{
			"SecretAccessKey": Schema{
				Type: ValueString,
			},
		},

		// AccessKeyId
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"Serial": Schema{
				Type: ValueNumber,
			},

			"Status": Schema{
				Type: EnumValue{
					Description: "Status",
					Options:     []string{"Active", "Inactive"},
				},
				Required: constraints.Always,
			},

			"UserName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},
		},
	}
}
