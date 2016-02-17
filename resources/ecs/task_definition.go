package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
func TaskDefinition() Resource {
	return Resource{
		AwsType: "AWS::ECS::TaskDefinition",

		// Name/ARN
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"ContainerDefinitions": Schema{
				Type:     containerDefinition,
				Array:    true,
				Required: constraints.Always,
			},

			"Volumes": Schema{
				Type:     volume,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
