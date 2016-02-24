package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html
var corsRule = NestedResource{
	Description: "S3 Cors Configuration Rule",
	Properties: Properties{
		"AllowedHeaders": Schema{
			Type: Multiple(ValueString),
		},

		"AllowedMethods": Schema{
			Type: Multiple(EnumValue{
				Description: "CORS Allowed Methods",
				Options:     []string{"GET", "PUT", "HEAD", "POST", "DELETE"},
			}),
			Required: constraints.Always,
		},

		"AllowedOrigins": Schema{
			Type:     Multiple(ValueString),
			Required: constraints.Always,
		},

		"ExposedHeaders": Schema{
			Type: Multiple(ValueString),
		},

		"Id": Schema{
			Type:         ValueString,
			ValidateFunc: StringLengthValidate(1, 255),
		},

		"MaxAge": Schema{
			Type: ValueNumber,
		},
	},
}
