package kinesis

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
var Stream = Resource{
	AwsType: "AWS::Kinesis::Stream",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ARN,
		},
	},

	// Physical ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"ShardCount": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
