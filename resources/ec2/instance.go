package ec2

import (
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

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
				Type:  networkInterface,
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
}
