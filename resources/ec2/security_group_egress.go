package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html
var SecurityGroupEgress = Resource{
	AwsType: "AWS::EC2::SecurityGroupEgress",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CidrIp": Schema{
			Type:      CIDR,
			Required:  constraints.PropertyNotExists("DestinationSecurityGroupId"),
			Conflicts: constraints.PropertyExists("DestinationSecurityGroupId"),
		},

		"DestinationSecurityGroupId": Schema{
			Type:      SecurityGroupID,
			Required:  constraints.PropertyNotExists("CidrIp"),
			Conflicts: constraints.PropertyExists("CidrIp"),
		},

		"FromPort": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"GroupId": Schema{
			Type:     SecurityGroupID,
			Required: constraints.Always,
		},

		"IpProtocol": Schema{
			Type:     ipProtocol,
			Required: constraints.Always,
		},

		"ToPort": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
