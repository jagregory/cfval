package s3

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
var lifecycleConfiguration = NestedResource{
	Description: "S3 Lifecycle Configuration",
	Properties: Properties{
		"Rules": Schema{
			Type:  lifecycleRule,
			Array: true,
		},
	},
}
