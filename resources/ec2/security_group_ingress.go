package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func SecurityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",

		Properties: Properties{
			"CidrIp": Schema{
				Type: CIDR,
			},

			"FromPort": Schema{
				Type:     ValueNumber,
				Required: constraints.Always,
			},

			"GroupId": Schema{
				Type: ValueString,
			},

			"GroupName": Schema{
				Type: ValueString,
			},

			"IpProtocol": Schema{
				Required: constraints.Always,
				Type:     ipProtocol,
			},

			"SourceSecurityGroupId": Schema{
				Type: ValueString,
			},

			"SourceSecurityGroupName": Schema{
				Type: ValueString,
			},

			"SourceSecurityGroupOwnerId": Schema{
				Type: ValueString,
			},

			"ToPort": Schema{
				Type:     ValueNumber,
				Required: constraints.Always,
			},
		},
	}
}
