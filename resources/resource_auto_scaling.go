package resources

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
var metricsCollection = NestedResource{
	Description: "Auto Scaling MetricsCollection",
	Properties: Properties{
		"Granularity": Schema{
			Type:     ValueString,
			Required: constraints.Always,
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
			Required: constraints.Always,
			Array:    true,
		},

		"TopicARN": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

var autoScalingTag = NestedResource{
	Description: "AutoScaling Tag",
	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
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
				Array:    true,
				Type:     AvailabilityZone,
				Required: constraints.PropertyNotExists("VPCZoneIdentifier"),
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
				Type:     ValueString,
				Required: constraints.PropertyNotExists("LaunchConfigurationName"),
			},

			"LaunchConfigurationName": Schema{
				Type:     ValueString,
				Required: constraints.PropertyNotExists("InstanceId"),
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
				Type:     ValueString,
				Array:    true,
				Required: constraints.PropertyNotExists("AvailabilityZones"),
			},
		},
	}
}

var volumeType = EnumValue{
	Description: "EBS Block Device Volume Type",

	Options: []string{"standard", "io1", "gp2"},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
var autoScalingEbsBlockDevice = NestedResource{
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
var autoScalingBlockDeviceMapping = NestedResource{
	Description: "AutoScaling Block Device Mapping",
	Properties: Properties{
		"DeviceName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Ebs": Schema{
			Type:      autoScalingEbsBlockDevice,
			Required:  constraints.PropertyNotExists("VirtualName"),
			Conflicts: constraints.PropertyExists("VirtualName"),
		},

		"NoDevice": Schema{
			Type: ValueBool,
		},

		"VirtualName": Schema{
			Type:      ValueString,
			Required:  constraints.PropertyNotExists("Ebs"),
			Conflicts: constraints.PropertyExists("Ebs"),
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
				Type:  autoScalingBlockDeviceMapping,
				Array: true,
			},

			"ClassicLinkVPCId": Schema{
				Type: ValueString,
			},

			"ClassicLinkVPCSecurityGroups": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.PropertyExists("ClassicLinkVPCId"),
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
				Required: constraints.Always,
			},

			"DefaultResult": Schema{
				Type: ValueString,
			},

			"HeartbeatTimeout": Schema{
				Type: ValueNumber,
			},

			"LifecycleTransition": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"NotificationMetadata": Schema{
				Type: ValueString,
			},

			// TODO: Do we need an ARN type?
			"NotificationTargetARN": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"RoleARN": Schema{
				Type:     ValueString,
				Required: constraints.Always,
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
			Required: constraints.Always,
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
				Required: constraints.Always,
			},

			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Cooldown": Schema{
				Type:      ValueString,
				Conflicts: constraints.PropertyNot("PolicyType", "StepScaling"),
			},

			"EstimatedInstanceWarmup": Schema{
				Type:      ValueNumber,
				Conflicts: constraints.PropertyNot("PolicyType", "StepScaling"),
			},

			"MetricAggregationType": Schema{
				Type:      metricAggregationType,
				Default:   "Average",
				Conflicts: constraints.PropertyNot("PolicyType", "StepScaling"),
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
				Type:      ValueNumber,
				Required:  constraints.PropertyIs("PolicyType", "SimpleScaling"),
				Conflicts: constraints.PropertyNot("PolicyType", "SimpleScaling"),
			},

			"StepAdjustments": Schema{
				Type:      stepAdjustment,
				Array:     true,
				Required:  constraints.PropertyIs("PolicyType", "StepScaling"),
				Conflicts: constraints.PropertyNot("PolicyType", "StepScaling"),
			},
		},
	}
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
func ScheduledAction() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::ScheduledAction",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AutoScalingGroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"DesiredCapacity": Schema{
				Type: ValueNumber,
			},

			"EndTime": Schema{
				Type: Timestamp,
			},

			"MaxSize": Schema{
				Type: ValueNumber,
			},

			"MinSize": Schema{
				Type: ValueNumber,
			},

			"Recurrence": Schema{
				Type: ValueString,
			},

			"StartTime": Schema{
				Type: Timestamp,
			},
		},
	}
}
