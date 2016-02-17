package sns

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
var Topic = Resource{
	AwsType: "AWS::SNS::Topic",

	// Topic ARN
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: map[string]Schema{
		"DisplayName": Schema{
			Type: ValueString,
		},

		"Subscription": Schema{
			Type:  snsSubscription,
			Array: true,
		},

		"TopicName": Schema{
			Type: ValueString,
		},
	},
}
