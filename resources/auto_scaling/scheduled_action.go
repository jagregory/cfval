package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
func ScheduledAction() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::ScheduledAction",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"DesiredCapacity": Schema{
				Type: ValueNumber,
			},

			"EndTime": Schema{
				Type: Timestamp,
			},

			"MaxSize": Schema{
				Type: ValueNumber,
			},

			"MinSize": Schema{
				Type: ValueNumber,
			},

			"Recurrence": Schema{
				Type: ValueString,
			},

			"StartTime": Schema{
				Type: Timestamp,
			},
		},
	}
}
