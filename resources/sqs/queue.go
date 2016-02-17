package sqs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
var Queue = Resource{
	AwsType: "AWS::SQS::Queue",

	Attributes: map[string]Schema{
		"Arn": Schema{
			Type: ValueString,
		},

		"QueueName": Schema{
			Type: ValueString,
		},
	},

	// Queue URL
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DelaySeconds": Schema{
			Type:         ValueNumber,
			Default:      0,
			ValidateFunc: IntegerRangeValidate(0, 900),
		},

		"MaximumMessageSize": Schema{
			Type:         ValueNumber,
			Default:      262144,
			ValidateFunc: IntegerRangeValidate(1024, 262144),
		},

		"MessageRetentionPeriod": Schema{
			Type:         ValueNumber,
			Default:      345600,
			ValidateFunc: IntegerRangeValidate(60, 1209600),
		},

		"QueueName": Schema{
			Type: ValueString,
		},

		"ReceiveMessageWaitTimeSeconds": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 20),
		},

		"RedrivePolicy": Schema{
			Type: redrivePolicy,
		},

		"VisibilityTimeout": Schema{
			Type:         ValueNumber,
			Default:      30,
			ValidateFunc: IntegerRangeValidate(0, 43200),
		},
	},
}
