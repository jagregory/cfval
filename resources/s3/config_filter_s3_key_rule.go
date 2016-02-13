package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html
var configFilterS3KeyRule = NestedResource{
	Description: "S3 NotificationConfiguration Config Filter S3Key Rule",
	Properties: Properties{
		"Name": Schema{
			Type: EnumValue{
				Description: "Name",
				Options:     []string{"prefix", "suffix"},
			},
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
