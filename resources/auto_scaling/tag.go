package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var autoScalingTag = NestedResource{
	Description: "AutoScaling Tag",
	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PropagateAtLaunch": Schema{
			Type: ValueBool,
		},
	},
}
