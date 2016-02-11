package resources

import . "github.com/jagregory/cfval/schema"

var resourceTag = NestedResource{
	Description: "Resource Tag",

	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: Always,
		},
	},
}
