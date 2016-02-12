package elasticache

import . "github.com/jagregory/cfval/schema"

var azMode = EnumValue{
	Description: "ElastiCache AZMode",

	Options: []string{"single-az", "cross-az"},
}
