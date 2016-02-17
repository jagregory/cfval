package s3

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
var Bucket = Resource{
	AwsType: "AWS::S3::Bucket",

	Attributes: map[string]Schema{
		"DomainName": Schema{
			Type: ValueString,
		},

		"WebsiteURL": Schema{
			Type: ValueString,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AccessControl": Schema{
			Type: accessControl,
		},

		"BucketName": Schema{
			Type: ValueString,
			ValidateFunc: RegexpValidate(
				`^[a-z0-9\.\-]+$`,
				"The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-).",
			),
		},

		"CorsConfiguration": Schema{
			Type: corsConfiguration,
		},

		"LifecycleConfiguration": Schema{
			Type: lifecycleConfiguration,
		},

		"LoggingConfiguration": Schema{
			Type: loggingConfiguration,
		},

		"NotificationConfiguration": Schema{
			Type: notificationConfiguration,
		},

		"ReplicationConfiguration": Schema{
			Type: replicationConfiguration,
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},

		"VersioningConfiguration": Schema{
			Type: versioningConfiguration,
		},

		"WebsiteConfiguration": Schema{
			Type: websiteConfiguration,
		},
	},
}
