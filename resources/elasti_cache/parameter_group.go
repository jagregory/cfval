package elasti_cache

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
var ParameterGroup = Resource{
	AwsType: "AWS::ElastiCache::ParameterGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CacheParameterGroupFamily": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Description": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Properties": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},
	},
}
