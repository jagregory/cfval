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
	AwsType:    "Auto Scaling MetricsCollection",
	Properties: map[string]Schema{
	// TODO
	},
}

var notificationConfiguration = Resource{
	AwsType:    "Auto Scaling NotificationConfiguration",
	Properties: map[string]Schema{
	// TODO
	},
}

func AutoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]Schema{
			"AvailabilityZones": ArrayOf(availabilityZone),

			"Cooldown": Schema{
				Type: TypeString,
			},

			"DesiredCapacity": Schema{
				Type: TypeString,
			},

			"HealthCheckGracePeriod": Schema{
				Type: TypeInteger,
			},

			"HealthCheckType": EnumOf("EC2", "ELB"),

			"InstanceId": Schema{
				Type: TypeString,
			},

			"LaunchConfigurationName": Schema{
				Type: TypeString,
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
				Type:  TypeString,
				Array: true,
			},
		},
	}
}

func LaunchConfiguration() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LaunchConfiguration",
		Properties: map[string]Schema{
			// "AssociatePublicIpAddress" : Boolean,
			// "BlockDeviceMappings" : [ BlockDeviceMapping, ... ],
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
