package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
var websiteConfigurationRoutingRule = NestedResource{
	Description: "S3 Website Configuration Routing Rule",
	Properties: Properties{
		"RedirectRule": Schema{
			Type:     websiteConfigurationRoutingRuleRedirectRule,
			Required: constraints.Always,
		},

		"RoutingRuleCondition": Schema{
			Type: websiteConfigurationRoutingRuleCondition,
		},
	},
}
