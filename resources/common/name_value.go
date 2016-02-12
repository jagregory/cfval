package common

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var NameValue = NestedResource{
	Description: "NameValue",

	Properties: Properties{
		"Name": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
