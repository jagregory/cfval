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
			Type:  ValueString,
			Array: true,
		},

		"AllowedMethods": Schema{
			Type: EnumValue{
				Description: "CORS Allowed Methods",
				Options:     []string{"GET", "PUT", "HEAD", "POST", "DELETE"},
			},
			Array:    true,
			Required: constraints.Always,
		},

		"AllowedOrigins": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.Always,
		},

		"ExposedHeaders": Schema{
			Type:  ValueString,
			Array: true,
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
