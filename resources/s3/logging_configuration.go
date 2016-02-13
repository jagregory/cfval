package s3

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
var loggingConfiguration = NestedResource{
	Description: "S3 Logging Configuration",
	Properties: Properties{
		"DestinationBucketName": Schema{
			Type: ValueString,
		},

		"LogFilePrefix": Schema{
			Type: ValueString,
		},
	},
}
