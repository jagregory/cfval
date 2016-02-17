package ec2

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internet-gateway.html
var InternetGateway = Resource{
	AwsType: "AWS::EC2::InternetGateway",

	// Name -- not sure about this. Docs say Name, but my testing we can Ref
	//         this into an InternetGatewayId property successfully.
	ReturnValue: Schema{
		Type: InternetGatewayID,
	},

	Properties: Properties{
		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},
	},
}
