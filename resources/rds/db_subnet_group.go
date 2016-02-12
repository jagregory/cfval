package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func DBSubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBSubnetGroup",
		Properties: map[string]Schema{
			"DBSubnetGroupDescription": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"SubnetIds": Schema{
				Type:     ValueString,
				Required: constraints.Always,
				Array:    true,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}
