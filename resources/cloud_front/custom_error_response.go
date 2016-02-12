package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html
var customErrorResponse = NestedResource{
	Description: "CloudFront DistributionConfig CustomErrorResponse",
	Properties: Properties{
		"ErrorCachingMinTTL": Schema{
			Type: ValueNumber,
		},

		"ErrorCode": Schema{
			Type:         ValueNumber,
			Required:     constraints.Always,
			ValidateFunc: NumberOptions(400, 403, 404, 405, 414, 500, 501, 502, 503, 504),
		},

		"ResponseCode": Schema{
			Type:         ValueNumber,
			Required:     constraints.PropertyExists("ResponsePagePath"),
			ValidateFunc: NumberOptions(200, 400, 403, 404, 405, 414, 500, 501, 502, 503, 504),
		},

		"ResponsePagePath": Schema{
			Type:     ValueString,
			Required: constraints.PropertyExists("ResponseCode"),
		},
	},
}
