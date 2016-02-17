package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html
var containerDefinitionVolumesFrom = NestedResource{
	Description: "EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom",

	Properties: Properties{
		"SourceContainer": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"ReadOnly": Schema{
			Type: ValueBool,
		},
	},
}
