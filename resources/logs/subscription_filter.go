package logs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
var SubscriptionFilter = Resource{
	AwsType: "AWS::Logs::SubscriptionFilter",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DestinationArn": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"FilterPattern": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"LogGroupName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"RoleArn": Schema{
			Type: ARN,
		},
	},
}
