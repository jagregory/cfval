package elasti_cache

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html
var SecurityGroup = Resource{
	AwsType: "AWS::ElastiCache::SecurityGroup",

	ReturnValue: Schema{
		Type: cacheSecurityGroupName,
	},

	Properties: Properties{
		"Description": Schema{
			Type: ValueString,
		},
	},
}
