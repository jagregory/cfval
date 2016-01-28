package main

func recordSet() Resource {
	return Resource{
		AwsType:    "AWS::Route53::RecordSet",
		Properties: map[string]Schema{},
	}
}
