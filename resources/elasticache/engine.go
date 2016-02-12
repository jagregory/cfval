package elasticache

import . "github.com/jagregory/cfval/schema"

var engine = EnumValue{
	Description: "ElastiCache Engine",

	Options: []string{"memcached", "redis"},
}
