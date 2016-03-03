package elasti_cache

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html
var SubnetGroup = Resource{
	AwsType: "AWS::ElastiCache::SubnetGroup",

	// Name
	ReturnValue: Schema{
		Type: cacheSubnetGroupName,
	},

	Properties: Properties{
		"Description": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SubnetIds": Schema{
			Type:     Multiple(SubnetID),
			Required: constraints.Always,
		},
	},
}
