package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html
var replicationConfigurationRule = NestedResource{
	Description: "S3 ReplicationConfiguration Rules",
	Properties: Properties{
		"Destination": Schema{
			Type:     replicationConfigurationRuleDestination,
			Required: constraints.Always,
		},

		"Id": Schema{
			Type: ValueString,
		},

		"Prefix": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Status": Schema{
			Type: EnumValue{
				Description: "Replication Rule Status",
				Options:     []string{"Enabled", "Disabled"},
			},
			Required: constraints.Always,
		},
	},
}
