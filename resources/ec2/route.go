package ec2

import . "github.com/jagregory/cfval/schema"

func Route() Resource {
	return Resource{
		AwsType: "AWS::EC2::Route",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"DestinationCidrBlock": Schema{
				Type:     CIDR,
				Required: Always,
			},

			"GatewayId": Schema{
				Type: ValueString,
			},

			"InstanceId": Schema{
				Type: ValueString,
			},

			"NetworkInterfaceId": Schema{
				Type: ValueString,
			},

			"RouteTableId": Schema{
				Type:     ValueString,
				Required: Always,
			},

			"VpcPeeringConnectionId": Schema{
				Type: ValueString,
			},
		},
	}
}
