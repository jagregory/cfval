package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
func LifecycleHook() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LifecycleHook",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"DefaultResult": Schema{
				Type: ValueString,
			},

			"HeartbeatTimeout": Schema{
				Type: ValueNumber,
			},

			"LifecycleTransition": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"NotificationMetadata": Schema{
				Type: ValueString,
			},

			// TODO: Do we need an ARN type?
			"NotificationTargetARN": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"RoleARN": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},
		},
	}
}
