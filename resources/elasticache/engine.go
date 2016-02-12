package elasticache

import . "github.com/jagregory/cfval/schema"

var engine = EnumValue{
	Description: "ElastiCache ReplicationGroup Engine",

	Options: []string{"redis"},
}
