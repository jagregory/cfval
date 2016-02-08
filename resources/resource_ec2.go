package resources

import . "github.com/jagregory/cfval/schema"

func Eip() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIP",

		// PublicIp
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"InstanceId": Schema{
				Type: ValueString,
			},

			"Domain": Schema{
				Type: ValueString,
			},
		},
	}
}

func Instance() Resource {
	return Resource{
		AwsType: "AWS::EC2::Instance",

		// InstanceId
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: ValueString,
			},

			// "BlockDeviceMappings":               ArrayOf(BlockDeviceMapping),
			// "DisableApiTermination":             Schema{Type: TypeBool},
			// "EbsOptimized":                      Schema{Type: TypeBool},
			// "IamInstanceProfile":                Schema{Type: TypeString},

			"ImageId": Schema{
				Type:     ValueString,
				Required: true,
			},

			// "InstanceInitiatedShutdownBehavior": Schema{Type: TypeString},

			"InstanceType": Schema{
				Type: ValueString,
			},

			// "KernelId":                          Schema{Type: TypeString},

			"KeyName": Schema{
				Type: ValueString,
			},

			// "Monitoring":                        Schema{Type: TypeBool},
			// "NetworkInterfaces":                 ArrayOf(NetworkInterface),
			// "PlacementGroupName":                Schema{Type: TypeString},
			// "PrivateIpAddress":                  Schema{Type: TypeString},
			// "RamdiskId":                         Schema{Type: TypeString},
			// "SecurityGroupIds":                  Schema{Type: TypeString, Array: true},
			// "SecurityGroups":                    Schema{Type: TypeString, Array: true},

			"SourceDestCheck": Schema{
				Type: ValueBool,
			},

			// "SsmAssociations":                   ArrayOf(SsmAssociation),

			"SubnetId": Schema{
				Type: ValueString,
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
		AwsType: "AWS::EC2::InternetGateway",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{},
	}
}

func Route() Resource {
	return Resource{
		AwsType: "AWS::EC2::Route",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"DestinationCidrBlock": Schema{
				Type:     cidr,
				Required: true,
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
				Required: true,
			},

			"VpcPeeringConnectionId": Schema{
				Type: ValueString,
			},
		},
	}
}

func RouteTable() Resource {
	return Resource{
		AwsType: "AWS::EC2::RouteTable",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"VpcId": Schema{
				Type:     vpcID,
				Required: true,
			},

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
							Required: true,
						},

						"IpProtocol": Schema{
							Type:     ValueString,
							Required: true,
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
							Required: true,
						},
					},
				},
			},
			"VpcId": Schema{Type: ValueString},
		},
	}
}

func SecurityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",

		Properties: Properties{
			"CidrIp": Schema{
				Type: cidr,
			},

			"FromPort": Schema{
				Type:     ValueNumber,
				Required: true,
			},

			"GroupId": Schema{
				Type: ValueString,
			},

			"GroupName": Schema{
				Type: ValueString,
			},

			"IpProtocol": Schema{
				Required: true,
				Type:     EnumValue{[]string{"tcp", "udp", "icmp", "-1"}},
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
				Required: true,
			},
		},
	}
}

func Subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",
		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: availabilityZone,
			},

			"CidrBlock": Schema{
				Type:     cidr,
				Required: true,
			},

			"MapPublicIpOnLaunch": Schema{
				Type: ValueBool,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"VpcId": Schema{
				Type:     vpcID,
				Required: true,
			},
		},
	}
}

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
				Required: true,
			},

			"SubnetId": Schema{
				Type:     ValueString,
				Required: true,
			},
		},
	}
}

func VpcGatewayAttachment() Resource {
	return Resource{
		AwsType: "AWS::EC2::VPCGatewayAttachment",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"InternetGatewayId": Schema{
				Type: ValueString,
			},

			"VpcId": Schema{
				Required: true,
				Type:     ValueString,
			},

			"VpnGatewayId": Schema{
				Type: ValueString,
			},
		},
	}
}
