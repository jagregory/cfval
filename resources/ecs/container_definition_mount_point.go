package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
var containerDefinitionMountPoint = NestedResource{
	Description: "EC2 Container Service TaskDefinition ContainerDefinitions MountPoints",

	Properties: Properties{
		"ContainerPath": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SourceVolume": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"ReadOnly": Schema{
			Type: ValueBool,
		},
	},
}
