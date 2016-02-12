package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html
var cookies = NestedResource{
	Description: "CloudFront ForwardedValues Cookies",

	Properties: Properties{
		"Forward": Schema{
			Type: EnumValue{
				Description: "CloudFront ForwardedValues Cookies Forward",
				Options:     []string{"none", "all", "whitelist"},
			},
			Required: constraints.Always,
		},

		"WhitelistedNames": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.PropertyIs("Forward", "whitelist"),
		},
	},
}
