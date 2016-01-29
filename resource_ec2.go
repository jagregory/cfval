package main

func eip() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIP",
		Properties: map[string]Schema{
			"InstanceId": Schema{Type: TypeString},
			"Domain":     Schema{Type: TypeString},
		},
	}
}

func instance() Resource {
	return Resource{
		AwsType: "AWS::EC2::Instance",
		Properties: map[string]Schema{
			"AvailabilityZone": Schema{Type: TypeString},
			// "BlockDeviceMappings":               ArrayOf(BlockDeviceMapping),
			// "DisableApiTermination":             Schema{Type: TypeBool},
			// "EbsOptimized":                      Schema{Type: TypeBool},
			// "IamInstanceProfile":                Schema{Type: TypeString},
			"ImageId": Schema{Type: TypeString, Required: true},
			// "InstanceInitiatedShutdownBehavior": Schema{Type: TypeString},
			"InstanceType": Schema{Type: TypeString},
			// "KernelId":                          Schema{Type: TypeString},
			"KeyName": Schema{Type: TypeString},
			// "Monitoring":                        Schema{Type: TypeBool},
			// "NetworkInterfaces":                 ArrayOf(NetworkInterface),
			// "PlacementGroupName":                Schema{Type: TypeString},
			// "PrivateIpAddress":                  Schema{Type: TypeString},
			// "RamdiskId":                         Schema{Type: TypeString},
			// "SecurityGroupIds":                  ArrayOf(Schema{Type: TypeString}),
			// "SecurityGroups":                    ArrayOf(Schema{Type: TypeString}),
			"SourceDestCheck": Schema{Type: TypeBool},
			// "SsmAssociations":                   ArrayOf(SsmAssociation),
			"SubnetId": Schema{Type: TypeString},
			"Tags":     ArrayOf(ResourceTag),
			// "Tenancy":                           Schema{Type: TypeString},
			// "UserData":                          Schema{Type: TypeString},
			// "Volumes":                           ArrayOf(MountPoint),
			// "AdditionalInfo":                    Schema{Type: TypeString},
		},
	}
}

func internetGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::InternetGateway",
		Properties: map[string]Schema{},
	}
}

func route() Resource {
	return Resource{
		AwsType: "AWS::EC2::Route",
		Properties: map[string]Schema{
			"DestinationCidrBlock":   Required(Cidr),
			"GatewayId":              Schema{Type: TypeString},
			"InstanceId":             Schema{Type: TypeString},
			"NetworkInterfaceId":     Schema{Type: TypeString},
			"RouteTableId":           Schema{Type: TypeString, Required: true},
			"VpcPeeringConnectionId": Schema{Type: TypeString},
		},
	}
}

func routeTable() Resource {
	return Resource{
		AwsType: "AWS::EC2::RouteTable",
		Properties: map[string]Schema{
			"VpcId": Required(VpcId),
			"Tags":  ArrayOf(ResourceTag),
		},
	}
}

func securityGroup() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroup",
		Properties: map[string]Schema{
			"GroupDescription": Schema{Type: TypeString},
			"SecurityGroupIngress": ArrayOf(Schema{
				Type: Resource{
					AwsType: "EC2 Security Group Rule Ingress",
					Properties: map[string]Schema{
						"CidrIp":                     Schema{Type: TypeString},
						"FromPort":                   Schema{Type: TypeInteger, Required: true},
						"IpProtocol":                 Schema{Type: TypeString, Required: true},
						"SourceSecurityGroupId":      Schema{Type: TypeString},
						"SourceSecurityGroupName":    Schema{Type: TypeString},
						"SourceSecurityGroupOwnerId": Schema{Type: TypeString},
						"ToPort":                     Schema{Type: TypeInteger, Required: true},
					},
				},
			}),
			"VpcId": Schema{Type: TypeString},
		},
	}
}

func securityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",
		Properties: map[string]Schema{
			"CidrIp":                     Cidr,
			"FromPort":                   Schema{Type: TypeInteger, Required: true},
			"GroupId":                    Schema{Type: TypeString},
			"GroupName":                  Schema{Type: TypeString},
			"IpProtocol":                 Required(EnumSchema("tcp", "udp", "icmp", "-1")),
			"SourceSecurityGroupId":      Schema{Type: TypeString},
			"SourceSecurityGroupName":    Schema{Type: TypeString},
			"SourceSecurityGroupOwnerId": Schema{Type: TypeString},
			"ToPort":                     Schema{Type: TypeInteger, Required: true},
		},
	}
}

func subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",
		Properties: map[string]Schema{
			"AvailabilityZone":    AvailabilityZone,
			"CidrBlock":           Required(Cidr),
			"MapPublicIpOnLaunch": Schema{Type: TypeBool},
			"Tags":                ArrayOf(ResourceTag),
			"VpcId":               Required(VpcId),
		},
	}
}

func subnetRouteTableAssociation() Resource {
	return Resource{
		AwsType: "AWS::EC2::SubnetRouteTableAssociation",
		Properties: map[string]Schema{
			"RouteTableId": Schema{Type: TypeString, Required: true},
			"SubnetId":     Schema{Type: TypeString, Required: true},
		},
	}
}

func vpcGatewayAttachment() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCGatewayAttachment",
		Properties: map[string]Schema{
			"InternetGatewayId": Schema{Type: TypeString},
			"VpcId":             Required(Schema{Type: TypeString}),
			"VpnGatewayId":      Schema{Type: TypeString},
		},
	}
}
