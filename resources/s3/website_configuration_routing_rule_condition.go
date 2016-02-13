package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html
var websiteConfigurationRoutingRuleCondition = NestedResource{
	Description: "S3 Website Configuration Routing Rules Routing Rule Condition",
	Properties: Properties{
		"HttpErrorCodeReturnedEquals": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("KeyPrefixEquals"),
		},

		"KeyPrefixEquals": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("HttpErrorCodeReturnedEquals"),
		},
	},
}
