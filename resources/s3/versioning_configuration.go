package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
var versioningConfiguration = NestedResource{
	Description: "S3 Versioning Configuration",
	Properties: Properties{
		"Status": Schema{
			Type: EnumValue{
				Description: "Versioning Status",
				Options:     []string{"Suspended", "Enabled"},
			},
			Required: constraints.Always,
		},
	},
}
