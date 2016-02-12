package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html
var cacheBehaviour = NestedResource{
	Description: "CloudFront DistributionConfig CacheBehavior",
	Properties: Properties{
		"AllowedMethods": Schema{
			Type: allowedMethods,
		},

		"CachedMethods": Schema{
			Type: cachedMethods,
		},

		"DefaultTTL": Schema{
			Type: ValueNumber,
		},

		"ForwardedValues": Schema{
			Type:     forwardedValues,
			Required: constraints.Always,
		},

		"MaxTTL": Schema{
			Type: ValueNumber,
		},

		"MinTTL": Schema{
			Type: ValueNumber,
		},

		"PathPattern": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SmoothStreaming": Schema{
			Type: ValueBool,
		},

		"TargetOriginId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TrustedSigners": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     viewerProtocolPolicy,
			Required: constraints.Always,
		},
	},
}
