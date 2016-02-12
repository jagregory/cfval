package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
var distributionConfig = NestedResource{
	Description: "CloudFront DistributionConfig",
	Properties: Properties{
		"Aliases": Schema{
			Type:  ValueString,
			Array: true,
		},

		"CacheBehaviors": Schema{
			Type:  cacheBehaviour,
			Array: true,
		},

		"Comment": Schema{
			Type: ValueString,
		},

		"CustomErrorResponses": Schema{
			Type:  customErrorResponse,
			Array: true,
		},

		"DefaultCacheBehavior": Schema{
			Required: constraints.Always,
			Type:     defaultCacheBehaviour,
		},

		"DefaultRootObject": Schema{
			Type: ValueString,
		},

		"Enabled": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"Logging": Schema{
			Type: logging,
		},

		"Origins": Schema{
			Array:    true,
			Required: constraints.Always,
			Type:     origin,
		},

		"PriceClass": Schema{
			Type: priceClass,
		},

		"Restrictions": Schema{
			Type: restrictions,
		},

		"ViewerCertificate": Schema{
			Type: viewerCertificate,
		},

		"WebACLId": Schema{
			Type: ValueString,
		},
	},
}
