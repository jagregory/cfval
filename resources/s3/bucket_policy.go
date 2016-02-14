package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
var bucketPolicy = Resource{
	AwsType: "AWS::S3::BucketPolicy",

	Properties: Properties{
		"Bucket": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},
	},
}
