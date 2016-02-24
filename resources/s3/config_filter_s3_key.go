package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
var configFilterS3Key = NestedResource{
	Description: "S3 NotificationConfiguration Config Filter S3Key",
	Properties: Properties{
		"Rules": Schema{
			Type:     Multiple(configFilterS3KeyRule),
			Required: constraints.Always,
		},
	},
}
