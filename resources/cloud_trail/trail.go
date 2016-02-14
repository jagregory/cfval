package cloud_trail

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
func Trail() Resource {
	return Resource{
		AwsType: "AWS::CloudTrail::Trail",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"CloudWatchLogsLogGroupArn": Schema{
				Type:     ValueString,
				Required: constraints.PropertyExists("CloudWatchLogsRoleArn"),
			},

			"CloudWatchLogsRoleArn": Schema{
				Type: ValueString,
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
				Type: ValueString,
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
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}