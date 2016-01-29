package main

func topic() Resource {
	return Resource{
		AwsType: "AWS::SNS::Topic",
		Properties: map[string]Schema{
			"DisplayName": Schema{Type: TypeString},
			// "Subscription": ArrayOf(Schema{
			// 	Type: Resource{...}
			// }),
			"TopicName": Schema{Type: TypeString},
		},
	}
}
