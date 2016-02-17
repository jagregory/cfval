package ecs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
func Cluster() Resource {
	return Resource{
		AwsType: "AWS::ECS::Cluster",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{},
	}
}
