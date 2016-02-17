package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html
var volume = NestedResource{
	Description: "EC2 Container Service TaskDefinition Volumes",

	Properties: Properties{
		"Name": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Host": Schema{
			Type: volumeHost,
		},
	},
}
