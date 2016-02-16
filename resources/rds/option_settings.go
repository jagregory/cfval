package rds

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html
var optionSettings = NestedResource{
	Description: "RDS OptionGroup OptionConfigurations OptionSetting",

	Properties: Properties{
		"Name": Schema{
			Type: ValueString,
		},

		"Value": Schema{
			Type: ValueString,
		},
	},
}
