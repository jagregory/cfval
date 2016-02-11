package resources

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
func Eip() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIP",

		// PublicIp
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"InstanceId": Schema{
				Type: InstanceID,
			},

			"Domain": Schema{
				Type: EnumValue{
					Description: "EIP Domain",
					Options:     []string{"vpc"},
				},
			},
		},
	}
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
var ec2EbsBlockDevice = NestedResource{
	Description: "Elastic Block Store Block Device",
	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"Encrypted": Schema{
			Type: ValueBool,
		},

		"Iops": Schema{
			Type:         ValueNumber,
			Required:     PropertyIs("VolumeType", "io1"),
			Conflicts:    PropertyNot("VolumeType", "io1"),
			ValidateFunc: IntegerRangeValidate(100, 2000),
		},

		"SnapshotId": Schema{
			Type: ValueString,
			// TODO: Required: Conditional If you specify both SnapshotId and VolumeSize, VolumeSize must be equal or greater than the size of the snapshot.
		},

		"VolumeSize": Schema{
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 1024),
			// TODO: If the volume type is io1, the minimum value is 10.
			// TODO: Required: Conditional If you specify both SnapshotId and VolumeSize, VolumeSize must be equal or greater than the size of the snapshot.
		},

		"VolumeType": Schema{
			Type: EnumValue{
				Description: "VolumeType",
				Options:     []string{"io1", "gp2"},
			},
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
var ec2BlockDeviceMapping = NestedResource{
	Description: "EC2 Block Device Mapping",
	Properties: Properties{
		"DeviceName": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"Ebs": Schema{
			Type:      ec2EbsBlockDevice,
			Required:  PropertyNotExists("VirtualName"),
			Conflicts: PropertyExists("VirtualName"),
		},

		"NoDevice": Schema{
			Type: JSON, // TODO: This should actually always be an empty map
		},

		"VirtualName": Schema{
			Type:      ValueString,
			Required:  PropertyNotExists("Ebs"),
			Conflicts: PropertyExists("Ebs"),
			ValidateFunc: RegexpValidate(
				"^ephemeral\\d+$",
				"The name must be in the form ephemeralX where X is a number starting from zero (0), for example, ephemeral0",
			),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html
var ec2InstanceType = EnumValue{
	Description: "EC2 Instance Type",
	Options: []string{
		// Current Generation Instances
		// General purpose
		"t2.nano", "t2.micro", "t2.small", "t2.medium", "t2.large", "m4.large", "m4.xlarge", "m4.2xlarge", "m4.4xlarge", "m4.10xlarge", "m3.medium", "m3.large", "m3.xlarge", "m3.2xlarge",
		// Compute optimized
		"c4.large", "c4.xlarge", "c4.2xlarge", "c4.4xlarge", "c4.8xlarge", "c3.large", "c3.xlarge", "c3.2xlarge", "c3.4xlarge", "c3.8xlarge",
		// Memory optimized
		"r3.large", "r3.xlarge", "r3.2xlarge", "r3.4xlarge", "r3.8xlarge",
		// Storage optimized
		"i2.xlarge", "i2.2xlarge", "i2.4xlarge", "i2.8xlarge", "d2.xlarge", "d2.2xlarge", "d2.4xlarge", "d2.8xlarge",
		// GPU instances
		"g2.2xlarge", "g2.8xlarge",

		// Previous Generation Instances
		// General purpose
		"m1.small", "m1.medium", "m1.large", "m1.xlarge",
		// Compute optimized
		"c1.medium", "c1.xlarge", "cc2.8xlarge",
		// Memory optimized
		"m2.xlarge", "m2.2xlarge", "m2.4xlarge", "cr1.8xlarge",
		// Storage optimized
		"hi1.4xlarge", "hs1.8xlarge",
		// GPU instances
		"cg1.4xlarge",
		// Micro instances
		"t1.micro",
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
var privateIPAddressSpecification = NestedResource{
	Description: "EC2 Network Interface Private IP Specification",
	Properties: Properties{
		"PrivateIpAddress": Schema{
			Type:     IPAddress,
			Required: Always,
		},

		"Primary": Schema{
			Type:     ValueBool,
			Required: Always,
			// TODO: You can set only one primary private IP address.
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
var ec2NetworkInterface = NestedResource{
	Description: "EC2 NetworkInterface",
	Properties: Properties{
		"AssociatePublicIpAddress": Schema{
			Type: ValueBool,
		},

		"DeleteOnTermination": Schema{
			Type: ValueBool,
		},

		"Description": Schema{
			Type: ValueString,
		},

		"DeviceIndex": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"GroupSet": Schema{
			Array: true,
			Type:  SecurityGroupID,
		},

		"NetworkInterfaceId": Schema{
			Type: ValueString,
		},

		"PrivateIpAddress": Schema{
			Type: IPAddress,
		},

		"PrivateIpAddresses": Schema{
			Array: true,
			Type:  privateIPAddressSpecification,
		},

		"SecondaryPrivateIpAddressCount": Schema{
			Type: ValueNumber,
		},

		"SubnetId": Schema{
			Type:     SubnetID,
			Required: PropertyNotExists("NetworkInterfaceId"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
var associationParameter = NestedResource{
	Description: "EC2 Instance SsmAssociations AssociationParameter",
	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"Value": Schema{
			Array:    true,
			Type:     ValueString,
			Required: Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
var ssmAssociation = NestedResource{
	Description: "EC2 Instance SsmAssociation",
	Properties: Properties{
		"AssociationParameters": Schema{
			Type:  associationParameter,
			Array: true,
		},

		"DocumentName": Schema{
			Type:     ValueString,
			Required: Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
var ec2MountPoint = NestedResource{
	Description: "EC2 MountPoint",
	Properties: Properties{
		"Device": Schema{
			Type:     ValueString,
			Required: Always,
		},

		"VolumeId": Schema{
			Type:     ValueString,
			Required: Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
func Instance() Resource {
	return Resource{
		AwsType: "AWS::EC2::Instance",

		// InstanceId
		ReturnValue: Schema{
			Type: InstanceID,
		},

		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: AvailabilityZone,
			},

			"BlockDeviceMappings": Schema{
				Type:  ec2BlockDeviceMapping,
				Array: true,
			},

			"DisableApiTermination": Schema{
				Type: ValueBool,
			},

			"EbsOptimized": Schema{
				Type:    ValueBool,
				Default: false,
			},

			"IamInstanceProfile": Schema{
				Type: ValueString,
			},

			"ImageId": Schema{
				Type:     ValueString,
				Required: Always,
			},

			"InstanceInitiatedShutdownBehavior": Schema{
				Type: EnumValue{
					Description: "Instance Shutdown Behaviour",
					Options:     []string{"stop", "terminate"},
				},
			},

			"InstanceType": Schema{
				Type: ec2InstanceType,
			},

			"KernelId": Schema{
				Type: ValueString,
			},

			"KeyName": Schema{
				Type: KeyName,
			},

			"Monitoring": Schema{
				Type: ValueBool,
			},

			"NetworkInterfaces": Schema{
				Array: true,
				Type:  ec2NetworkInterface,
				Conflicts: Constraints{
					PropertyExists("SecurityGroupIds"),
					PropertyExists("SubnetId"),
				},
			},

			"PlacementGroupName": Schema{
				Type: ValueString,
			},

			"PrivateIpAddress": Schema{
				Type: IPAddress,
			},

			"RamdiskId": Schema{
				Type: ValueString,
			},

			"SecurityGroupIds": Schema{
				Type:      ValueString,
				Array:     true,
				Conflicts: PropertyExists("NetworkInterfaces"),
			},

			"SecurityGroups": Schema{
				Type:  SecurityGroupName,
				Array: true,
			},

			"SourceDestCheck": Schema{
				Type: ValueBool,
			},

			"SsmAssociations": Schema{
				Type:  ssmAssociation,
				Array: true,
			},

			"SubnetId": Schema{
				Type:      SubnetID,
				Conflicts: PropertyExists("NetworkInterfaces"),
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"Tenancy": Schema{
				Type: EnumValue{
					Description: "EC2 Instance Tenancy",
					Options:     []string{"default", "dedicated"},
				},
			},

			"UserData": Schema{
				Type: ValueString,
			},

			"Volumes": Schema{
				Type:  ec2MountPoint,
				Array: true,
			},

			"AdditionalInfo": Schema{
				Type: ValueString,
			},
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
				Type:     CIDR,
				Required: Always,
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
				Required: Always,
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
				Type:     VpcID,
				Required: Always,
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
							Required: Always,
						},

						"IpProtocol": Schema{
							Type:     ValueString,
							Required: Always,
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
							Required: Always,
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

var ipProtocol = EnumValue{
	Description: "SecurityGroupIngress IpProtocol",

	Options: []string{"tcp", "udp", "icmp", "-1"},
}

func SecurityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",

		Properties: Properties{
			"CidrIp": Schema{
				Type: CIDR,
			},

			"FromPort": Schema{
				Type:     ValueNumber,
				Required: Always,
			},

			"GroupId": Schema{
				Type: ValueString,
			},

			"GroupName": Schema{
				Type: ValueString,
			},

			"IpProtocol": Schema{
				Required: Always,
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
				Required: Always,
			},
		},
	}
}

func Subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",
		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: AvailabilityZone,
			},

			"CidrBlock": Schema{
				Type:     CIDR,
				Required: Always,
			},

			"MapPublicIpOnLaunch": Schema{
				Type: ValueBool,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"VpcId": Schema{
				Type:     VpcID,
				Required: Always,
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
				Required: Always,
			},

			"SubnetId": Schema{
				Type:     ValueString,
				Required: Always,
			},
		},
	}
}

func VPCGatewayAttachment() Resource {
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
				Required: Always,
				Type:     ValueString,
			},

			"VpnGatewayId": Schema{
				Type: ValueString,
			},
		},
	}
}
