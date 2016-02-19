package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
var LifecycleHook = Resource{
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

		"NotificationTargetARN": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"RoleARN": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},
	},
}
