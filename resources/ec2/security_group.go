package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html
func SecurityGroup() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroup",

		Attributes: map[string]Schema{
			"GroupId": Schema{
				Type: SecurityGroupID,
			},
		},

		// SecurityGroupName for non-VPC, SecurityGroupId for VPC
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"GroupDescription": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"SecurityGroupEgress": Schema{
				Array: true,
				Type: NestedResource{
					Description: "EC2 Security Group Rule Egress",
					Properties: Properties{
						"CidrIp": Schema{
							Type:      CIDR,
							Conflicts: constraints.PropertyExists("DestinationSecurityGroupId"),
						},

						"FromPort": Schema{
							Type:     ValueNumber,
							Required: constraints.Always,
						},

						"IpProtocol": Schema{
							Type:     ipProtocol,
							Required: constraints.Always,
						},

						"DestinationSecurityGroupId": Schema{
							Type:      SecurityGroupID,
							Conflicts: constraints.PropertyExists("CidrIp"),
						},

						"ToPort": Schema{
							Type:     ValueNumber,
							Required: constraints.Always,
						},
					},
				},
			},

			"SecurityGroupIngress": Schema{
				Array: true,
				Type: NestedResource{
					Description: "EC2 Security Group Rule Ingress",
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

						"IpProtocol": Schema{
							Type:     ipProtocol,
							Required: constraints.Always,
						},

						"SourceSecurityGroupId": Schema{
							Type:      SecurityGroupID,
							Conflicts: constraints.PropertyExists("CidrIp"),
						},

						"SourceSecurityGroupName": Schema{
							Type:      SecurityGroupName,
							Conflicts: constraints.PropertyExists("CidrIp"),
						},

						// TODO: This is an AWS Account ID
						"SourceSecurityGroupOwnerId": Schema{
							Type: ValueString,
						},

						"ToPort": Schema{
							Type:     ValueNumber,
							Required: constraints.Always,
						},
					},
				},
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VpcId": Schema{
				Type: VpcID,
			},
		},
	}
}
