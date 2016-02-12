package elasticache

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func SubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"SubnetIds": Schema{
				Type:     ValueString,
				Required: constraints.Always,
				Array:    true,
			},
		},
	}
}
