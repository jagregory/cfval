package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/deprecations"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
var distributionConfig = NestedResource{
	Description: "CloudFront DistributionConfig",
	Properties: Properties{
		"Aliases": Schema{
			Type: Multiple(ValueString),
		},

		"CacheBehaviors": Schema{
			Type: Multiple(cacheBehaviour),
		},

		"Comment": Schema{
			Type: ValueString,
		},

		"CNAMEs": Schema{
			Deprecated: deprecations.Deprecated("CNAMEs should be specified in the Aliases property."),
		},

		"CustomErrorResponses": Schema{
			Type: Multiple(customErrorResponse),
		},

		"CustomOrigin": Schema{
			Deprecated: deprecations.Deprecated("An CustomOrigin should be specified as an item in the Origins property."),
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
			Type:     Multiple(origin),
			Required: constraints.Always,
		},

		"PriceClass": Schema{
			Type: priceClass,
		},

		"Restrictions": Schema{
			Type: restrictions,
		},

		"S3Origin": Schema{
			Deprecated: deprecations.Deprecated("An S3Origin should be specified as an item in the Origins property."),
		},

		"ViewerCertificate": Schema{
			Type: viewerCertificate,
		},

		"WebACLId": Schema{
			Type: ValueString,
		},
	},
}
