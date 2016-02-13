package s3

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html
var notificationConfiguration = NestedResource{
	Description: "S3 NotificationConfiguration",
	Properties: Properties{
		"LambdaConfigurations": Schema{
			Type:  lambdaConfiguration,
			Array: true,
		},

		"QueueConfigurations": Schema{
			Type:  queueConfiguration,
			Array: true,
		},

		"TopicConfigurations": Schema{
			Type:  topicConfiguration,
			Array: true,
		},
	},
}
