package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
var TaskDefinition = Resource{
	AwsType: "AWS::ECS::TaskDefinition",

	// Name/ARN
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"ContainerDefinitions": Schema{
			Type:     Multiple(containerDefinition),
			Required: constraints.Always,
		},

		"Volumes": Schema{
			Type:     Multiple(volume),
			Required: constraints.Always,
		},
	},
}
