package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var notificationConfiguration = NestedResource{
	Description: "Auto Scaling NotificationConfiguration",
	Properties: Properties{
		"NotificationTypes": Schema{
			Type:     autoScalingNotificationType,
			Required: constraints.Always,
			Array:    true,
		},

		"TopicARN": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
