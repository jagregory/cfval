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
			Type: Multiple(ValueString),
		},

		"Cpu": Schema{
			Type: ValueNumber,
		},

		"EntryPoint": Schema{
			Type: Multiple(ValueString),
		},

		"Environment": Schema{
			Type: Multiple(containerDefinitionEnvironment),
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
			Type: Multiple(ValueString),
		},

		"Memory": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"MountPoints": Schema{
			Type: Multiple(containerDefinitionMountPoint),
		},

		"Name": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PortMappings": Schema{
			Type: Multiple(containerDefinitionPortMapping),
		},

		"VolumesFrom": Schema{
			Type: Multiple(containerDefinitionVolumesFrom),
		},
	},
}
