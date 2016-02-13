package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html
var replicationConfigurationRuleDestination = NestedResource{
	Description: "S3 ReplicationConfiguration Rules Destination",
	Properties: Properties{
		"Bucket": Schema{
			Type:     ValueString, // TODO: ARN
			Required: constraints.Always,
		},

		"StorageClass": Schema{
			Type: storageClass,
		},
	},
}
