package ecs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
var volumeHost = NestedResource{
	Description: "EC2 Container Service TaskDefinition Volumes Host",

	Properties: Properties{
		"SourcePath": Schema{
			Type: ValueString,
		},
	},
}
