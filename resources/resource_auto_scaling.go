package resources

import . "github.com/jagregory/cfval/schema"

var autoScalingTag = Resource{
	AwsType: "AutoScaling Tag",
	Properties: map[string]Schema{
		"Key": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Value": Schema{
			Type:     TypeString,
			Required: true,
		},

		"PropagateAtLaunch": Schema{
			Type: TypeBool,
		},
	},
}

var metricsCollection = Resource{
	AwsType: "Auto Scaling MetricsCollection",
	Properties: map[string]Schema{
		"Granularity": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Metrics": Schema{
			Type:  TypeString,
			Array: true,
		},
	},
}

var notificationConfiguration = Resource{
	AwsType: "Auto Scaling NotificationConfiguration",
	Properties: map[string]Schema{
		"NotificationTypes": Schema{
			Type:         TypeString,
			Required:     true,
			Array:        true,
			ValidateFunc: EnumValidate("autoscaling:EC2_INSTANCE_LAUNCH", "autoscaling:EC2_INSTANCE_LAUNCH_ERROR", "autoscaling:EC2_INSTANCE_TERMINATE", "autoscaling:EC2_INSTANCE_TERMINATE_ERROR", "autoscaling:TEST_NOTIFICATION"),
		},

		"TopicARN": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
func AutoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]Schema{
			"AvailabilityZones": Schema{
				Array:          true,
				Type:           TypeString,
				ValidateFunc:   availabilityZone,
				RequiredUnless: []string{"VPCZoneIdentifier"},
			},

			"Cooldown": Schema{
				Type: TypeString,
			},

			"DesiredCapacity": Schema{
				Type: TypeString,
			},

			"HealthCheckGracePeriod": Schema{
				Type: TypeInteger,
			},

			"HealthCheckType": Schema{
				Type:         TypeString,
				ValidateFunc: EnumValidate("EC2", "ELB"),
			},

			"InstanceId": Schema{
				Type:           TypeString,
				RequiredUnless: []string{"LaunchConfigurationName"},
			},

			"LaunchConfigurationName": Schema{
				Type:           TypeString,
				RequiredUnless: []string{"InstanceId"},
			},

			"LoadBalancerNames": Schema{
				Type:  TypeString,
				Array: true,
			},

			"MaxSize": Schema{
				Type: TypeString,
			},

			"MetricsCollection": Schema{
				Type:  metricsCollection,
				Array: true,
			},

			"MinSize": Schema{
				Type: TypeString,
			},

			"NotificationConfigurations": Schema{
				Type:  notificationConfiguration,
				Array: true,
			},

			"PlacementGroup": Schema{
				Type: TypeString,
			},

			"Tags": Schema{
				Type:  autoScalingTag,
				Array: true,
			},

			"TerminationPolicies": Schema{
				Type:  TypeString,
				Array: true,
			},

			"VPCZoneIdentifier": Schema{
				Type:           TypeString,
				Array:          true,
				RequiredUnless: []string{"AvailabilityZones"},
			},
		},
	}
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
var ebsBlockDevice = Resource{
	AwsType: "AutoScaling EBS Block Device",
	Properties: map[string]Schema{
		"DeleteOnTermination": Schema{
			Type:    TypeBool,
			Default: true,
		},

		"Encrypted": Schema{
			Type: TypeBool,
		},

		"Iops": Schema{
			Type: TypeInteger,
		},

		"SnapshotId": Schema{
			Type: TypeString,
		},

		"VolumeSize": Schema{
			Type:         TypeInteger,
			ValidateFunc: IntegerRangeValidate(1, 1024),
		},

		"VolumeType": Schema{
			Type:         TypeString,
			Default:      "standard",
			ValidateFunc: EnumValidate("standard", "io1", "gp2"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
var blockDeviceMapping = Resource{
	AwsType: "AutoScaling Block Device Mapping",
	Properties: map[string]Schema{
		"DeviceName": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Ebs": Schema{
			Type:           ebsBlockDevice,
			RequiredUnless: []string{"VirtualName"},
		},

		"NoDevice": Schema{
			Type: TypeBool,
		},

		"VirtualName": Schema{
			Type:           TypeString,
			RequiredUnless: []string{"Ebs"},
			ValidateFunc: RegexpValidate(
				"^ephemeral\\d+$",
				"The name must be in the form ephemeralX where X is a number starting from zero (0), for example, ephemeral0",
			),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
func LaunchConfiguration() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LaunchConfiguration",
		Properties: map[string]Schema{
			"AssociatePublicIpAddress": Schema{
				Type: TypeBool,
			},

			"BlockDeviceMappings": Schema{
				Type:  blockDeviceMapping,
				Array: true,
			},
			// "ClassicLinkVPCId" : String,
			// "ClassicLinkVPCSecurityGroups" : [ String, ... ],
			// "EbsOptimized" : Boolean,

			"IamInstanceProfile": Schema{
				Type: TypeString,
			},

			"ImageId": Schema{
				Type: TypeString,
			},

			// "InstanceId" : String,
			// "InstanceMonitoring" : Boolean,

			"InstanceType": Schema{
				Type: TypeString,
			},

			// "KernelId" : String,
			"KeyName": Schema{
				Type: TypeString,
			},

			// "PlacementTenancy" : String,
			// "RamDiskId" : String,

			"SecurityGroups": Schema{
				Type:  TypeString,
				Array: true,
			},

			// "SpotPrice" : String,

			"UserData": Schema{
				Type: TypeString,
			},
		},
	}
}
