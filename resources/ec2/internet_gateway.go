package ec2

import . "github.com/jagregory/cfval/schema"

func InternetGateway() Resource {
	return Resource{
		AwsType: "AWS::EC2::InternetGateway",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{},
	}
}
