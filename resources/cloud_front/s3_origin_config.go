package cloud_front

import . "github.com/jagregory/cfval/schema"

var s3OriginConfig = NestedResource{
	Description: "CloudFront DistributionConfig Origin S3Origin",
	Properties: Properties{
		"OriginAccessIdentity": Schema{
			Type: ValueString,
		},
	},
}
