package sqs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
var QueuePolicy = Resource{
	AwsType: "AWS::SQS::QueuePolicy",

	Properties: Properties{
		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"Queues": Schema{
			Type:     Multiple(ValueString),
			Required: constraints.Always,
		},
	},
}
