package main

func policy() Resource {
	return Resource{
		AwsType:    "AWS::IAM::Policy",
		Properties: map[string]Schema{},
	}
}

func role() Resource {
	return Resource{
		AwsType:    "AWS::IAM::Role",
		Properties: map[string]Schema{},
	}
}
