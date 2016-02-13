package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html
var lifecycleRule = NestedResource{
	Description: "AWS::S3::LifecycleRule",
	Properties: Properties{
		"ExpirationDate": Schema{
			Type: ValueString,
			Required: constraints.All{
				constraints.PropertyNotExists("ExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionTransition"),
				constraints.PropertyNotExists("Transition"),
			},
		},

		"ExpirationInDays": Schema{
			Type: ValueNumber,
			Required: constraints.All{
				constraints.PropertyNotExists("ExpirationDate"),
				constraints.PropertyNotExists("NoncurrentVersionExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionTransition"),
				constraints.PropertyNotExists("Transition"),
			},
		},

		"Id": Schema{
			Type:         ValueString,
			ValidateFunc: StringLengthValidate(1, 255),
		},

		"NoncurrentVersionExpirationInDays": Schema{
			Type: ValueNumber,
			Required: constraints.All{
				constraints.PropertyNotExists("ExpirationDate"),
				constraints.PropertyNotExists("ExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionTransition"),
				constraints.PropertyNotExists("Transition"),
			},
		},

		"NoncurrentVersionTransition": Schema{
			Type: noncurrentVersionTransition,
			Required: constraints.All{
				constraints.PropertyNotExists("ExpirationDate"),
				constraints.PropertyNotExists("ExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionExpirationInDays"),
				constraints.PropertyNotExists("Transition"),
			},
		},

		"Prefix": Schema{
			Type: ValueString,
		},

		"Status": Schema{
			Type: EnumValue{
				Description: "S3 Status",
				Options:     []string{"Enabled", "Disabled"},
			},
			Required: constraints.Always,
		},

		"Transition": Schema{
			Type: lifecycleRuleTransition,
			Required: constraints.All{
				constraints.PropertyNotExists("ExpirationDate"),
				constraints.PropertyNotExists("ExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionExpirationInDays"),
				constraints.PropertyNotExists("NoncurrentVersionTransition"),
			},
		},
	},
}
