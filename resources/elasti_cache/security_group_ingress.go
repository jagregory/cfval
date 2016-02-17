package elasti_cache

import . "github.com/jagregory/cfval/schema"
import "github.com/jagregory/cfval/constraints"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
var SecurityGroupIngress = Resource{
	AwsType: "AWS::ElastiCache::SecurityGroupIngress",

	Properties: Properties{
		"CacheSecurityGroupName": Schema{
			Type:     cacheSecurityGroupName,
			Required: constraints.Always,
		},

		"EC2SecurityGroupName": Schema{
			Type:     SecurityGroupName,
			Required: constraints.Always,
		},

		"EC2SecurityGroupOwnerId": Schema{
			Type: ValueString, // TODO: Account ID
		},
	},
}
