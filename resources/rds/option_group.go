package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
var OptionGroup = Resource{
	AwsType: "AWS::RDS::OptionGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"EngineName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"MajorEngineVersion": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OptionGroupDescription": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OptionConfigurations": Schema{
			Type:     Multiple(optionConfiguration),
			Required: constraints.Always,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
