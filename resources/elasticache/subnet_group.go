package elasticache

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html
func SubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"SubnetIds": Schema{
				Type:     SubnetID,
				Required: constraints.Always,
				Array:    true,
			},
		},
	}
}
