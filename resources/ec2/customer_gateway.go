package ec2

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)
import "github.com/jagregory/cfval/constraints"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html
func CustomerGateway() Resource {
	return Resource{
		AwsType: "AWS::EC2::CustomerGateway",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"BgpAsn": Schema{
				Type:     ValueNumber,
				Required: constraints.Always,
			},

			"IpAddress": Schema{
				Type:     IPAddress,
				Required: constraints.Always,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"Type": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},
		},
	}
}
