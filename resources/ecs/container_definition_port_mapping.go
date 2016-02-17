package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html
var containerDefinitionPortMapping = NestedResource{
	Description: "EC2 Container Service TaskDefinition ContainerDefinitions PortMapping",

	Properties: Properties{
		"ContainerPort": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"HostPort": Schema{
			Type: ValueNumber,
			// TODO: Do not specify a host port in the 49153 to 65535 port range;
			//       these ports are reserved for automatic assignment. Other reserved
			//       ports include 22 for SSH, the Docker ports 2375 and 2376, and the
			//       Amazon EC2 Container Service container agent port 51678. In
			//       addition, do not specify a host port that is being used for a
			//       task; that port is reserved while the task is running.
		},
	},
}
