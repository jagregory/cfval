package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html
var queueConfiguration = NestedResource{
	Description: "Simple Storage Service NotificationConfiguration QueueConfiguration",
	Properties: Properties{
		"Event": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Filter": Schema{
			Type: configFilter,
		},

		"Queue": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
