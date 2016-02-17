package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
var containerDefinition = NestedResource{
	Description: "EC2 Container Service TaskDefinition ContainerDefinition",

	Properties: Properties{
		"Command": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Cpu": Schema{
			Type: ValueNumber,
		},

		"EntryPoint": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Environment": Schema{
			Type:  containerDefinitionEnvironment,
			Array: true,
		},

		"Essential": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"Image": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Links": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Memory": Schema{
			Type:  ValueString,
			Array: true,
		},

		"MountPoints": Schema{
			Type:     containerDefinitionMountPoint,
			Array:    true,
			Required: constraints.Always,
		},

		"Name": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PortMappings": Schema{
			Type:  containerDefinitionPortMapping,
			Array: true,
		},

		"VolumesFrom": Schema{
			Type:  containerDefinitionVolumesFrom,
			Array: true,
		},
	},
}
