package ec2

import . "github.com/jagregory/cfval/schema"

func SubnetRouteTableAssociation() Resource {
	return Resource{
		AwsType: "AWS::EC2::SubnetRouteTableAssociation",

		// ID
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"RouteTableId": Schema{
				Type:     ValueString,
				Required: Always,
			},

			"SubnetId": Schema{
				Type:     ValueString,
				Required: Always,
			},
		},
	}
}
