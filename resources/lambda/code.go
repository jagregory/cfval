package lambda

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html
var code = NestedResource{
	Description: "Lambda Function Code",

	Properties: Properties{
		"S3Bucket": Schema{
			Type: ValueString,
			Required: constraints.Any{
				constraints.PropertyExists("S3Key"),
				constraints.PropertyNotExists("ZipFile"),
			},
		},

		"S3Key": Schema{
			Type: ValueString,
			Required: constraints.Any{
				constraints.PropertyExists("S3Bucket"),
				constraints.PropertyNotExists("ZipFile"),
			},
		},

		"S3ObjectVersion": Schema{
			Type: ValueString,
			Conflicts: constraints.Any{
				constraints.PropertyNotExists("S3Bucket"),
				constraints.PropertyNotExists("S3Key"),
			},
		},

		"ZipFile": Schema{
			Type: ValueString,
			Required: constraints.All{
				constraints.PropertyNotExists("S3Bucket"),
				constraints.PropertyNotExists("S3Key"),
			},
		},
	},
}
