package rds

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html
var optionConfiguration = NestedResource{
	Description: "RDS OptionGroup OptionConfiguration",

	Properties: Properties{
		"DBSecurityGroupMemberships": Schema{
			Type:      dbSecurityGroupName,
			Array:     true,
			Conflicts: constraints.PropertyExists("VPCSecurityGroupMemberships"),
		},

		"OptionName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OptionSettings": Schema{
			Type: optionSettings,
		},

		"Port": Schema{
			Type: ValueNumber,
		},

		"VpcSecurityGroupMemberships": Schema{
			Type:      SecurityGroupID,
			Array:     true,
			Conflicts: constraints.PropertyExists("DBSecurityGroupMemberships"),
		},
	},
}
