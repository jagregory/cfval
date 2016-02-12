package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var optionsSettings = NestedResource{
	Description: "Elastic Beanstalk OptionSettings",
	Properties: Properties{
		"Namespace": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OptionName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
