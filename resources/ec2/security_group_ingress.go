package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
func SecurityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",

		Properties: Properties{
			"CidrIp": Schema{
				Type: CIDR,
				Conflicts: constraints.Any{
					constraints.PropertyExists("SourceSecurityGroupName"),
					constraints.PropertyExists("SourceSecurityGroupId"),
				},
			},

			"FromPort": Schema{
				Type:     ValueNumber,
				Required: constraints.Always,
			},

			"GroupId": Schema{
				Type:     SecurityGroupID,
				Required: constraints.PropertyNotExists("GroupName"),
			},

			"GroupName": Schema{
				Type:     SecurityGroupName,
				Required: constraints.PropertyNotExists("GroupId"),
			},

			"IpProtocol": Schema{
				Required: constraints.Always,
				Type:     ipProtocol,
			},

			"SourceSecurityGroupId": Schema{
				Type:      SecurityGroupID,
				Conflicts: constraints.PropertyExists("CidrIp"),
			},

			"SourceSecurityGroupName": Schema{
				Type:      SecurityGroupName,
				Conflicts: constraints.PropertyExists("CidrIp"),
			},

			// TODO: AWS account type
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
