package ec2

import . "github.com/jagregory/cfval/schema"

func VPCGatewayAttachment() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCGatewayAttachment",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"InternetGatewayId": Schema{
				Type: ValueString,
			},

			"VpcId": Schema{
				Required: Always,
				Type:     ValueString,
			},

			"VpnGatewayId": Schema{
				Type: ValueString,
			},
		},
	}
}
