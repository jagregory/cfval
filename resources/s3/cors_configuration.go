package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
var corsConfiguration = NestedResource{
	Description: "S3 Cors Configuration",
	Properties: Properties{
		"CorsRules": Schema{
			Type:     Multiple(corsRule),
			Required: constraints.Always,
		},
	},
}
