package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
var networkInterface = NestedResource{
	Description: "EC2 Instance NetworkInterface",

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
			Required: constraints.Always,
		},

		"GroupSet": Schema{
			Array: true,
			Type:  SecurityGroupID,
		},

		"NetworkInterfaceId": Schema{
			Type: NetworkInterfaceID,
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
			Required: constraints.PropertyNotExists("NetworkInterfaceId"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
var Instance = Resource{
	AwsType: "AWS::EC2::Instance",

	Attributes: map[string]Schema{
		"AvailabilityZone": Schema{
			Type: AvailabilityZone,
		},

		"PrivateDnsName": Schema{
			Type: ValueString,
		},

		"PublicDnsName": Schema{
			Type: ValueString,
		},

		"PrivateIp": Schema{
			Type: IPAddress,
		},

		"PublicIp": Schema{
			Type: IPAddress,
		},
	},

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
			Type:     ImageID,
			Required: constraints.Always,
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
			Type:  networkInterface,
			Conflicts: constraints.Any{
				constraints.PropertyExists("SecurityGroupIds"),
				constraints.PropertyExists("SubnetId"),
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
			Type:      SecurityGroupID,
			Array:     true,
			Conflicts: constraints.PropertyExists("NetworkInterfaces"),
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
			Conflicts: constraints.PropertyExists("NetworkInterfaces"),
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
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
			Type:  mountPoint,
			Array: true,
		},

		"AdditionalInfo": Schema{
			Type: ValueString,
		},
	},
}
