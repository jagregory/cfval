package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
var configFilter = NestedResource{
	Description: "S3 NotificationConfiguration Config Filter",
	Properties: Properties{
		"S3Key": Schema{
			Type:     configFilterS3Key,
			Required: constraints.Always,
		},
	},
}
