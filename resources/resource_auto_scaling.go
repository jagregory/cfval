package resources

import . "github.com/jagregory/cfval/schema"

var metricsCollection = NestedResource{
	Description: "Auto Scaling MetricsCollection",
	Properties: Properties{
		"Granularity": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Metrics": Schema{
			Type:  ValueString,
			Array: true,
		},
	},
}

var notificationConfiguration = NestedResource{
	Description: "Auto Scaling NotificationConfiguration",
	Properties: Properties{
		"NotificationTypes": Schema{
			Type:     EnumValue{[]string{"autoscaling:EC2_INSTANCE_LAUNCH", "autoscaling:EC2_INSTANCE_LAUNCH_ERROR", "autoscaling:EC2_INSTANCE_TERMINATE", "autoscaling:EC2_INSTANCE_TERMINATE_ERROR", "autoscaling:TEST_NOTIFICATION"}},
			Required: true,
			Array:    true,
		},

		"TopicARN": Schema{
			Type:     ValueString,
			Required: true,
		},
	},
}

var autoScalingTag = NestedResource{
	Description: "AutoScaling Tag",
	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: true,
		},

		"PropagateAtLaunch": Schema{
			Type: ValueBool,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
func AutoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AvailabilityZones": Schema{
				Array:          true,
				Type:           availabilityZone,
				RequiredUnless: []string{"VPCZoneIdentifier"},
			},

			"Cooldown": Schema{
				Type: ValueString,
			},

			"DesiredCapacity": Schema{
				Type: ValueString,
			},

			"HealthCheckGracePeriod": Schema{
				Type: ValueNumber,
			},

			"HealthCheckType": Schema{
				Type: EnumValue{[]string{"EC2", "ELB"}},
			},

			"InstanceId": Schema{
				Type:           ValueString,
				RequiredUnless: []string{"LaunchConfigurationName"},
			},

			"LaunchConfigurationName": Schema{
				Type:           ValueString,
				RequiredUnless: []string{"InstanceId"},
			},

			"LoadBalancerNames": Schema{
				Type:  ValueString,
				Array: true,
			},

			"MaxSize": Schema{
				Type: ValueString,
			},

			"MetricsCollection": Schema{
				Type:  metricsCollection,
				Array: true,
			},

			"MinSize": Schema{
				Type: ValueString,
			},

			"NotificationConfigurations": Schema{
				Type:  notificationConfiguration,
				Array: true,
			},

			"PlacementGroup": Schema{
				Type: ValueString,
			},

			"Tags": Schema{
				Type:  autoScalingTag,
				Array: true,
			},

			"TerminationPolicies": Schema{
				Type:  ValueString,
				Array: true,
			},

			"VPCZoneIdentifier": Schema{
				Type:           ValueString,
				Array:          true,
				RequiredUnless: []string{"AvailabilityZones"},
			},
		},
	}
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
var ebsBlockDevice = NestedResource{
	Description: "AutoScaling EBS Block Device",
	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"Encrypted": Schema{
			Type: ValueBool,
		},

		"Iops": Schema{
			Type: ValueNumber,
		},

		"SnapshotId": Schema{
			Type: ValueString,
		},

		"VolumeSize": Schema{
			Type: IntegerRangeValidate(1, 1024),
		},

		"VolumeType": Schema{
			Type:    EnumValue{[]string{"standard", "io1", "gp2"}},
			Default: "standard",
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
var blockDeviceMapping = NestedResource{
	Description: "AutoScaling Block Device Mapping",
	Properties: Properties{
		"DeviceName": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Ebs": Schema{
			Type:           ebsBlockDevice,
			RequiredUnless: []string{"VirtualName"},
		},

		"NoDevice": Schema{
			Type: ValueBool,
		},

		"VirtualName": Schema{
			Type: FuncType(RegexpValidate(
				"^ephemeral\\d+$",
				"The name must be in the form ephemeralX where X is a number starting from zero (0), for example, ephemeral0",
			)),
			RequiredUnless: []string{"Ebs"},
		},
	},
}

var iamInstanceProfileType = StringLengthValidate(1, 1600)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
func LaunchConfiguration() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LaunchConfiguration",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AssociatePublicIpAddress": Schema{
				Type: ValueBool,
			},

			"BlockDeviceMappings": Schema{
				Type:  blockDeviceMapping,
				Array: true,
			},

			"ClassicLinkVPCId": Schema{
				Type: ValueString,
			},

			"ClassicLinkVPCSecurityGroups": Schema{
				Type:       ValueString,
				Array:      true,
				RequiredIf: []string{"ClassicLinkVPCId"},
			},

			"EbsOptimized": Schema{
				Type:    ValueBool,
				Default: false,
			},

			"IamInstanceProfile": Schema{
				Type: iamInstanceProfileType,
			},

			"ImageId": Schema{
				Type: ValueString,
			},

			"InstanceId": Schema{
				Type: ValueString,
			},

			"InstanceMonitoring": Schema{
				Type:    ValueBool,
				Default: true,
			},

			"InstanceType": Schema{
				Type: ValueString,
			},

			"KernelId": Schema{
				Type: ValueString,
			},

			"KeyName": Schema{
				Type: ValueString,
			},

			// TODO: If you specify this property, you must specify at least one subnet in the VPCZoneIdentifier property of the AWS::AutoScaling::AutoScalingGroup resource.
			"PlacementTenancy": Schema{
				Type: EnumValue{[]string{"default", "dedicated"}},
			},

			"RamDiskId": Schema{
				Type: ValueString,
			},

			"SecurityGroups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"SpotPrice": Schema{
				Type: ValueString,
			},

			"UserData": Schema{
				Type: ValueString,
			},
		},
	}
}
