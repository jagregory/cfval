package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
var websiteConfiguration = NestedResource{
	Description: "S3 Website Configuration",
	Properties: Properties{
		"ErrorDocument": Schema{
			Type: ValueString,
		},

		"IndexDocument": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"RedirectAllRequestsTo": Schema{
			Type: websiteConfigurationRedirectAllRequestsTo,
		},

		"RoutingRules": Schema{
			Type: Multiple(websiteConfigurationRoutingRule),
		},
	},
}
