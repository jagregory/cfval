package sns

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
var Topic = Resource{
	AwsType: "AWS::SNS::Topic",

	ReturnValue: Schema{
		Type: ARN,
	},

	Properties: Properties{
		"DisplayName": Schema{
			Type: ValueString,
		},

		"Subscription": Schema{
			Type: Multiple(snsSubscription),
		},

		"TopicName": Schema{
			Type: ValueString,
		},
	},
}
