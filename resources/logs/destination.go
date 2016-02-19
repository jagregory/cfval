package logs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
var Destination = Resource{
	AwsType: "AWS::Logs::Destination",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DestinationName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"DestinationPolicy": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"RoleArn": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"TargetArn": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},
	},
}
