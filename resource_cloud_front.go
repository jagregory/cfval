package main

func distribution() Resource {
	return Resource{
		AwsType:    "AWS::CloudFront::Distribution",
		Properties: map[string]Schema{},
	}
}
