package resources

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
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

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
var autoScalingNotificationType = EnumValue{
	Description: "Auto Scaling Notification Type",

	Options: []string{
		"autoscaling:EC2_INSTANCE_LAUNCH",
		"autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
		"autoscaling:EC2_INSTANCE_TERMINATE",
		"autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
		"autoscaling:TEST_NOTIFICATION",
	},
}

var notificationConfiguration = NestedResource{
	Description: "Auto Scaling NotificationConfiguration",
	Properties: Properties{
		"NotificationTypes": Schema{
			Type:     autoScalingNotificationType,
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

var healthCheckType = EnumValue{
	Description: "Auto Scaling Health Check Type",

	Options: []string{"EC2", "ELB"},
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
				Type:           AvailabilityZone,
				RequiredUnless: Constraints{PropertyExists("VPCZoneIdentifier")},
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
				Type: healthCheckType,
			},

			"InstanceId": Schema{
				Type:           ValueString,
				RequiredUnless: Constraints{PropertyExists("LaunchConfigurationName")},
			},

			"LaunchConfigurationName": Schema{
				Type:           ValueString,
				RequiredUnless: Constraints{PropertyExists("InstanceId")},
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
				RequiredUnless: Constraints{PropertyExists("AvailabilityZones")},
			},
		},
	}
}

var volumeType = EnumValue{
	Description: "EBS Block Device Volume Type",

	Options: []string{"standard", "io1", "gp2"},
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
			Type:         ValueNumber,
			ValidateFunc: IntegerRangeValidate(1, 1024),
		},

		"VolumeType": Schema{
			Type:    volumeType,
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
			RequiredUnless: Constraints{PropertyExists("VirtualName")},
		},

		"NoDevice": Schema{
			Type: ValueBool,
		},

		"VirtualName": Schema{
			Type:           ValueString,
			RequiredUnless: Constraints{PropertyExists("Ebs")},
			ValidateFunc: RegexpValidate(
				"^ephemeral\\d+$",
				"The name must be in the form ephemeralX where X is a number starting from zero (0), for example, ephemeral0",
			),
		},
	},
}

var placementTenancy = EnumValue{
	Description: "Launch Configuration Placement Tenancy",

	Options: []string{"default", "dedicated"},
}

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
				RequiredIf: Constraints{PropertyExists("ClassicLinkVPCId")},
			},

			"EbsOptimized": Schema{
				Type:    ValueBool,
				Default: false,
			},

			"IamInstanceProfile": Schema{
				Type:         ValueString,
				ValidateFunc: StringLengthValidate(1, 1600),
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
				Type: KeyName,
			},

			// TODO: If you specify this property, you must specify at least one subnet in the VPCZoneIdentifier property of the AWS::AutoScaling::AutoScalingGroup resource.
			// This will require some reverse lookups from this resource to any which use it: not supported yet.
			"PlacementTenancy": Schema{
				Type: placementTenancy,
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

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
func LifecycleHook() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LifecycleHook",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: true,
			},

			"DefaultResult": Schema{
				Type: ValueString,
			},

			"HeartbeatTimeout": Schema{
				Type: ValueNumber,
			},

			"LifecycleTransition": Schema{
				Type:     ValueString,
				Required: true,
			},

			"NotificationMetadata": Schema{
				Type: ValueString,
			},

			// TODO: Do we need an ARN type?
			"NotificationTargetARN": Schema{
				Type:     ValueString,
				Required: true,
			},

			"RoleARN": Schema{
				Type:     ValueString,
				Required: true,
			},
		},
	}
}

var metricAggregationType = EnumValue{
	Description: "ScalingPolicy MetricAggregationType",

	Options: []string{"Minimum", "Maximum", "Average"},
}

var policyType = EnumValue{
	Description: "ScalingPolicy PolicyType",

	Options: []string{"SimpleScaling", "StepScaling"},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html
var stepAdjustment = NestedResource{
	Description: "Auto Scaling ScalingPolicy StepAdjustments",

	Properties: Properties{
		"MetricIntervalLowerBound": Schema{
			Type: ValueNumber,
		},

		"MetricIntervalUpperBound": Schema{
			Type: ValueNumber,
		},

		"ScalingAdjustment": Schema{
			Type:     ValueNumber,
			Required: true,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
func ScalingPolicy() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::ScalingPolicy",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AdjustmentType": Schema{
				Type:     ValueString,
				Required: true,
			},

			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: true,
			},

			"Cooldown": Schema{
				Type:      ValueString,
				Conflicts: Constraints{PropertyNot("PolicyType", "StepScaling")},
			},

			"EstimatedInstanceWarmup": Schema{
				Type:      ValueNumber,
				Conflicts: Constraints{PropertyNot("PolicyType", "StepScaling")},
			},

			"MetricAggregationType": Schema{
				Type:      metricAggregationType,
				Default:   "Average",
				Conflicts: Constraints{PropertyNot("PolicyType", "StepScaling")},
			},

			// TODO: This property replaces the MinAdjustmentStep property
			"MinAdjustmentMagnitude": Schema{
				Type: ValueNumber,
			},

			"PolicyType": Schema{
				Type:    policyType,
				Default: "SimpleScaling",
			},

			"ScalingAdjustment": Schema{
				Type:       ValueNumber,
				RequiredIf: Constraints{PropertyIs("PolicyType", "SimpleScaling")},  // Required: Conditional. This property is required if the policy type is SimpleScaling.
				Conflicts:  Constraints{PropertyNot("PolicyType", "SimpleScaling")}, // This property is not supported with any other policy type.
			},

			"StepAdjustments": Schema{
				Type:       stepAdjustment,
				Array:      true,
				RequiredIf: Constraints{PropertyIs("PolicyType", "StepScaling")},  // Required: Conditional. This property is required if the policy type is StepScaling.
				Conflicts:  Constraints{PropertyNot("PolicyType", "StepScaling")}, // This property is not supported with any other policy type.
			},
		},
	}
}
