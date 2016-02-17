package ecs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html
var loadBalancer = NestedResource{
	Description: "EC2 Container Service Service LoadBalancer",

	Properties: Properties{
		"ContainerName": Schema{
			Type: ValueString,
		},

		"ContainerPort": Schema{
			Type: ValueNumber,
		},

		"LoadBalancerName": Schema{
			Type: ValueString,
		},
	},
}
