package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html
var lambdaConfiguration = NestedResource{
	Description: "Simple Storage Service NotificationConfiguration LambdaConfiguration",
	Properties: Properties{
		"Event": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Filter": Schema{
			Type: configFilter,
		},

		"Function": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
