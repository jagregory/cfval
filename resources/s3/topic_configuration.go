package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
var topicConfiguration = NestedResource{
	Description: "S3 NotificationConfiguration TopicConfiguration",
	Properties: Properties{
		"Event": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Filter": Schema{
			Type: configFilter,
		},

		"Topic": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
