package sns

import . "github.com/jagregory/cfval/schema"

func Topic() Resource {
	return Resource{
		AwsType: "AWS::SNS::Topic",

		// Topic ARN
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"DisplayName": Schema{
				Type: ValueString,
			},

			// "Subscription": ArrayOf(Schema{
			// 	Type: Resource{...}
			// }),

			"TopicName": Schema{
				Type: ValueString,
			},
		},
	}
}
