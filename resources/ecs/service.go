package ecs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
var Service = Resource{
	AwsType: "AWS::ECS::Service",

	// ARN
	ReturnValue: Schema{
		Type: ValueString,
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
			Type:  loadBalancer,
			Array: true,
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
