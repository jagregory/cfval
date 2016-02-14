package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
func VPCDHCPOptionsAssociation() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCDHCPOptionsAssociation",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"DhcpOptionsId": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"VpcId": Schema{
				Type:     VpcID,
				Required: constraints.Always,
			},
		},
	}
}
