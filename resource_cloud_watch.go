package main

func alarm() Resource {
	return Resource{
		AwsType:    "AWS::CloudWatch::Alarm",
		Properties: map[string]Schema{},
	}
}
