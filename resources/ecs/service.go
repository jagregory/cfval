package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
var Service = Resource{
	AwsType: "AWS::ECS::Service",

	ReturnValue: Schema{
		Type: ARN,
	},

	Properties: Properties{
		"Cluster": Schema{
			Type: ValueString,
		},

		"DesiredCount": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"LoadBalancers": Schema{
			Type: Multiple(loadBalancer),
		},

		"Role": Schema{
			Type:     ValueString,
			Required: constraints.PropertyExists("LoadBalancers"),
		},

		"TaskDefinition": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
