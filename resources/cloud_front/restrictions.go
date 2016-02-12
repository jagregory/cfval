package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
var restrictions = NestedResource{
	Description: "CloudFront DistributionConfiguration Restrictions",
	Properties: Properties{
		"GeoRestriction": Schema{
			Type:     geoRestriction,
			Required: constraints.Always,
		},
	},
}
