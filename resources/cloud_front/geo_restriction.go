package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html
var geoRestriction = NestedResource{
	Description: "CloudFront DistributionConfig Restrictions GeoRestriction",
	Properties: Properties{
		"Locations": Schema{
			Type:  common.CountryCode,
			Array: true,
			Required: constraints.Any{
				constraints.PropertyIs("RestrictionType", "blacklist"),
				constraints.PropertyIs("RestrictionType", "whitelist"),
			},
		},

		"RestrictionType": Schema{
			Type: EnumValue{
				Description: "RestrictionType",
				Options:     []string{"blacklist", "whitelist", "none"},
			},
			Required: constraints.Always,
		},
	},
}
