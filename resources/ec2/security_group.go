package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func SecurityGroup() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroup",

		// SecurityGroupName for non-VPC, SecurityGroupId for VPC
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"GroupDescription": Schema{
				Type: ValueString,
			},

			"SecurityGroupIngress": Schema{
				Array: true,
				Type: NestedResource{
					Description: "EC2 Security Group Rule Ingress",
					Properties: Properties{
						"CidrIp": Schema{
							Type: ValueString,
						},

						"FromPort": Schema{
							Type:     ValueNumber,
							Required: constraints.Always,
						},

						"IpProtocol": Schema{
							Type:     ValueString,
							Required: constraints.Always,
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
				},
			},

			"VpcId": Schema{
				Type: VpcID,
			},
		},
	}
}
