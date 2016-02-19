package sqs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues-redrivepolicy.html
var redrivePolicy = NestedResource{
	Description: "SQS RedrivePolicy",

	Properties: Properties{
		"deadLetterTargetArn": Schema{
			Type: ARN,
		},

		"maxReceiveCount": Schema{
			Type: ValueNumber,
		},
	},
}
