package elasti_cache

import . "github.com/jagregory/cfval/schema"

var cacheNodeType = EnumValue{
	Description: "ElastiCache NodeType",

	Options: []string{
		"cache.t2.micro", "cache.t2.small", "cache.t2.medium",
		"cache.m3.medium", "cache.m3.large", "cache.m3.xlarge", "cache.m3.2xlarge",
		"cache.t1.micro", "cache.m1.small", "cache.m1.medium", "cache.m1.large",
		"cache.m1.xlarge", "cache.c1.xlarge", "cache.r3.large", "cache.r3.xlarge",
		"cache.r3.2xlarge", "cache.r3.4xlarge", "cache.r3.8xlarge", "cache.m2.xlarge",
		"cache.m2.2xlarge", "cache.m2.4xlarge",
	},
}
