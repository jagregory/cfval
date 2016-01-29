package resources

import . "github.com/jagregory/cfval/schema"

func Eip() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIP",
		Properties: map[string]Schema{
			"InstanceId": Schema{
				Type: TypeString,
			},

			"Domain": Schema{
				Type: TypeString,
			},
		},
	}
}

func Instance() Resource {
	return Resource{
		AwsType: "AWS::EC2::Instance",
		Properties: map[string]Schema{
			"AvailabilityZone": Schema{
				Type: TypeString,
			},

			// "BlockDeviceMappings":               ArrayOf(BlockDeviceMapping),
			// "DisableApiTermination":             Schema{Type: TypeBool},
			// "EbsOptimized":                      Schema{Type: TypeBool},
			// "IamInstanceProfile":                Schema{Type: TypeString},

			"ImageId": Schema{
				Type:     TypeString,
				Required: true,
			},

			// "InstanceInitiatedShutdownBehavior": Schema{Type: TypeString},

			"InstanceType": Schema{
				Type: TypeString,
			},

			// "KernelId":                          Schema{Type: TypeString},

			"KeyName": Schema{
				Type: TypeString,
			},

			// "Monitoring":                        Schema{Type: TypeBool},
			// "NetworkInterfaces":                 ArrayOf(NetworkInterface),
			// "PlacementGroupName":                Schema{Type: TypeString},
			// "PrivateIpAddress":                  Schema{Type: TypeString},
			// "RamdiskId":                         Schema{Type: TypeString},
			// "SecurityGroupIds":                  Schema{Type: TypeString, Array: true},
			// "SecurityGroups":                    Schema{Type: TypeString, Array: true},

			"SourceDestCheck": Schema{
				Type: TypeBool,
			},

			// "SsmAssociations":                   ArrayOf(SsmAssociation),

			"SubnetId": Schema{
				Type: TypeString,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			// "Tenancy":                           Schema{Type: TypeString},
			// "UserData":                          Schema{Type: TypeString},
			// "Volumes":                           ArrayOf(MountPoint),
			// "AdditionalInfo":                    Schema{Type: TypeString},
		},
	}
}

func InternetGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::InternetGateway",
		Properties: map[string]Schema{},
	}
}

func Route() Resource {
	return Resource{
		AwsType: "AWS::EC2::Route",
		Properties: map[string]Schema{
			"DestinationCidrBlock": Required(cidr),

			"GatewayId": Schema{
				Type: TypeString,
			},

			"InstanceId": Schema{
				Type: TypeString,
			},

			"NetworkInterfaceId": Schema{
				Type: TypeString,
			},

			"RouteTableId": Schema{
				Type:     TypeString,
				Required: true,
			},

			"VpcPeeringConnectionId": Schema{
				Type: TypeString,
			},
		},
	}
}

func RouteTable() Resource {
	return Resource{
		AwsType: "AWS::EC2::RouteTable",
		Properties: map[string]Schema{
			"VpcId": Required(vpcId),

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},
		},
	}
}

func SecurityGroup() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroup",
		Properties: map[string]Schema{
			"GroupDescription": Schema{
				Type: TypeString,
			},

			"SecurityGroupIngress": Schema{
				Array: true,
				Type: Resource{
					AwsType: "EC2 Security Group Rule Ingress",
					Properties: map[string]Schema{
						"CidrIp": Schema{
							Type: TypeString,
						},

						"FromPort": Schema{
							Type:     TypeInteger,
							Required: true,
						},

						"IpProtocol": Schema{
							Type:     TypeString,
							Required: true,
						},

						"SourceSecurityGroupId": Schema{
							Type: TypeString,
						},

						"SourceSecurityGroupName": Schema{
							Type: TypeString,
						},

						"SourceSecurityGroupOwnerId": Schema{
							Type: TypeString,
						},

						"ToPort": Schema{
							Type:     TypeInteger,
							Required: true,
						},
					},
				},
			},
			"VpcId": Schema{Type: TypeString},
		},
	}
}

func SecurityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",
		Properties: map[string]Schema{
			"CidrIp": cidr,

			"FromPort": Schema{
				Type:     TypeInteger,
				Required: true,
			},

			"GroupId": Schema{
				Type: TypeString,
			},

			"GroupName": Schema{
				Type: TypeString,
			},

			"IpProtocol": Required(EnumOf("tcp", "udp", "icmp", "-1")),

			"SourceSecurityGroupId": Schema{
				Type: TypeString,
			},

			"SourceSecurityGroupName": Schema{
				Type: TypeString,
			},

			"SourceSecurityGroupOwnerId": Schema{
				Type: TypeString,
			},

			"ToPort": Schema{
				Type:     TypeInteger,
				Required: true,
			},
		},
	}
}

func Subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",
		Properties: map[string]Schema{
			"AvailabilityZone": availabilityZone,

			"CidrBlock": Required(cidr),

			"MapPublicIpOnLaunch": Schema{
				Type: TypeBool,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"VpcId": Required(vpcId),
		},
	}
}

func SubnetRouteTableAssociation() Resource {
	return Resource{
		AwsType: "AWS::EC2::SubnetRouteTableAssociation",
		Properties: map[string]Schema{
			"RouteTableId": Schema{
				Type:     TypeString,
				Required: true,
			},

			"SubnetId": Schema{
				Type:     TypeString,
				Required: true,
			},
		},
	}
}

func VpcGatewayAttachment() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCGatewayAttachment",
		Properties: map[string]Schema{
			"InternetGatewayId": Schema{
				Type: TypeString,
			},

			"VpcId": Required(Schema{
				Type: TypeString,
			}),

			"VpnGatewayId": Schema{
				Type: TypeString,
			},
		},
	}
}
