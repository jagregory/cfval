package rds

import "github.com/jagregory/cfval/schema"

var engine = schema.EnumValue{
	Description: "DB Instance Engine",
	Options: []string{
		"MySQL",
		"mariadb",
		"oracle-se1",
		"oracle-se",
		"oracle-ee",
		"sqlserver-ee",
		"sqlserver-se",
		"sqlserver-ex",
		"sqlserver-web",
		"postgres",
		"aurora",
	},
}
