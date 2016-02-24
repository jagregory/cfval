package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var defaultCacheBehaviour = NestedResource{
	Description: "CloudFront DefaultCacheBehaviour",
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
			Required: constraints.Always,
			Type:     forwardedValues,
		},

		"MaxTTL": Schema{
			Type: ValueNumber,
		},

		"MinTTL": Schema{
			Type: ValueString,
		},

		"SmoothStreaming": Schema{
			Type: ValueBool,
		},

		"TargetOriginId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TrustedSigners": Schema{
			Type: Multiple(ValueString),
		},

		"ViewerProtocolPolicy": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
