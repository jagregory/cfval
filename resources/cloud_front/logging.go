package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
var logging = NestedResource{
	Description: "CloudFront Logging",
	Properties: Properties{
		"Bucket": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"IncludeCookies": Schema{
			Type: ValueBool,
		},

		"Prefix": Schema{
			Type: ValueString,
		},
	},
}
