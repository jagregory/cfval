package s3

import . "github.com/jagregory/cfval/schema"

var accessControl = EnumValue{
	Description: "S3 Bucket AccessControl",

	Options: []string{"AuthenticatedRead", "AwsExecRead", "BucketOwnerRead", "BucketOwnerFullControl", "LogDeliveryWrite", "Private", "PublicRead", "PublicReadWrite"},
}
