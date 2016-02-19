package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html
var replicationConfiguration = NestedResource{
	Description: "S3 ReplicationConfiguration",
	Properties: Properties{
		"Role": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"Rules": Schema{
			Type:     replicationConfigurationRule,
			Required: constraints.Always,
		},
	},
}
