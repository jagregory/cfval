package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html
var forwardedValues = NestedResource{
	Description: "CloudFront ForwardedValues",
	Properties: Properties{
		"Cookies": Schema{
			Type: cookies,
		},

		"Headers": Schema{
			Type:  ValueString,
			Array: true,
		},

		"QueryString": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},
	},
}
