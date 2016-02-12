package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
var customOriginConfig = NestedResource{
	Description: "CloudFront DistributionConfig Origin CustomOrigin",
	Properties: Properties{
		"HTTPPort": Schema{
			Type: ValueString,
		},

		"HTTPSPort": Schema{
			Type: ValueString,
		},

		"OriginProtocolPolicy": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
