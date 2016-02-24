package cloud_trail

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
var Trail = Resource{
	AwsType: "AWS::CloudTrail::Trail",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CloudWatchLogsLogGroupArn": Schema{
			Type:     ARN,
			Required: constraints.PropertyExists("CloudWatchLogsRoleArn"),
		},

		"CloudWatchLogsRoleArn": Schema{
			Type: ARN,
		},

		"EnableLogFileValidation": Schema{
			Type: ValueBool,
		},

		"IncludeGlobalServiceEvents": Schema{
			Type: ValueBool,
		},

		"IsLogging": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"KMSKeyId": Schema{
			Type: ARN,
		},

		"S3BucketName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"S3KeyPrefix": Schema{
			Type: ValueString,
		},

		"SnsTopicName": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
