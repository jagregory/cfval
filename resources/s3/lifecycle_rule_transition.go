package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html
var lifecycleRuleTransition = NestedResource{
	Description: "S3 Lifecycle Rule Transition",
	Properties: Properties{
		"StorageClass": Schema{
			Type:     storageClass,
			Required: constraints.Always,
		},

		"TransitionDate": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("TransitionInDays"),
		},

		"TransitionInDays": Schema{
			Type:     ValueNumber,
			Required: constraints.PropertyNotExists("TransitionDate"),
		},
	},
}
