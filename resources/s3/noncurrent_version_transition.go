package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html
var noncurrentVersionTransition = NestedResource{
	Description: "S3 Lifecycle Rule NoncurrentVersionTransition",
	Properties: Properties{
		"StorageClass": Schema{
			Type:     storageClass,
			Required: constraints.Always,
		},

		"TransitionInDays": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
