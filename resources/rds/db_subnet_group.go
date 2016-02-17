package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
var DBSubnetGroup = Resource{
	AwsType: "AWS::RDS::DBSubnetGroup",
	Properties: map[string]Schema{
		"DBSubnetGroupDescription": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SubnetIds": Schema{
			Type:     SubnetID,
			Required: constraints.Always,
			Array:    true,
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},
	},
}
