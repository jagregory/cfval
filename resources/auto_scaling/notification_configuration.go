package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var notificationConfiguration = NestedResource{
	Description: "Auto Scaling NotificationConfiguration",
	Properties: Properties{
		"NotificationTypes": Schema{
			Type:     Multiple(autoScalingNotificationType),
			Required: constraints.Always,
		},

		"TopicARN": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},
	},
}
