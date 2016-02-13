package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html
var websiteConfigurationRoutingRuleRedirectRule = NestedResource{
	Description: "S3 Website Configuration Routing Rules Redirect Rule",
	Properties: Properties{
		"HostName": Schema{
			Type: ValueString,
		},

		"HttpRedirectCode": Schema{
			Type: ValueString,
		},

		"Protocol": Schema{
			Type: ValueString,
		},

		"ReplaceKeyPrefixWith": Schema{
			Type:      ValueString,
			Conflicts: constraints.PropertyExists("ReplaceKeyWith"),
		},

		"ReplaceKeyWith": Schema{
			Type:      ValueString,
			Conflicts: constraints.PropertyExists("ReplaceKeyPrefixWith"),
		},
	},
}
