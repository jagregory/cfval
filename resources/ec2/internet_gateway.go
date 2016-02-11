package ec2

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internet-gateway.html
func InternetGateway() Resource {
	return Resource{
		AwsType: "AWS::EC2::InternetGateway",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}
