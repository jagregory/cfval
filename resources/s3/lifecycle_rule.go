package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var s3LifecycleRule = NestedResource{
	Description: "AWS::S3::LifecycleRule",
	Properties: Properties{
		// "ExpirationDate":   Schema{Type: TypeString},
		"ExpirationInDays": Schema{
			Type: ValueNumber,
		},

		"Id": Schema{
			Type: ValueString,
		},

		// "NoncurrentVersionExpirationInDays": Schema{Type: TypeInteger},
		// "NoncurrentVersionTransition":       S3LifecycleRuleNoncurrentVersionTransition,
		// "Prefix":                            Schema{Type: TypeString},

		"Status": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		// "Transition":                        S3LifecycleRuleTransition,
	},
}
