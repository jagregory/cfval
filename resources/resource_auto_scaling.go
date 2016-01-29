package resources

import . "github.com/jagregory/cfval/schema"

var autoScalingTag = Schema{
	Type: Resource{
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
	},
}

func AutoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]Schema{
			"AvailabilityZones":          ArrayOf(availabilityZone),
			"Cooldown":                   Schema{Type: TypeString},
			"DesiredCapacity":            Schema{Type: TypeString},
			"HealthCheckGracePeriod":     Schema{Type: TypeInteger},
			"HealthCheckType":            EnumSchema("EC2", "ELB"),
			"InstanceId":                 Schema{Type: TypeString},
			"LaunchConfigurationName":    Schema{Type: TypeString},
			"LoadBalancerNames":          ArrayOf(Schema{Type: TypeString}),
			"MaxSize":                    Schema{Type: TypeString},
			"MetricsCollection":          ArrayOf(metricsCollection),
			"MinSize":                    Schema{Type: TypeString},
			"NotificationConfigurations": ArrayOf(notificationConfiguration),
			"PlacementGroup":             Schema{Type: TypeString},
			"Tags":                       ArrayOf(autoScalingTag),
			"TerminationPolicies":        ArrayOf(Schema{Type: TypeString}),
			"VPCZoneIdentifier":          ArrayOf(Schema{Type: TypeString}),
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
			"IamInstanceProfile": Schema{Type: TypeString},
			"ImageId":            Schema{Type: TypeString},
			// "InstanceId" : String,
			// "InstanceMonitoring" : Boolean,
			"InstanceType": Schema{Type: TypeString},
			// "KernelId" : String,
			"KeyName": Schema{Type: TypeString},
			// "PlacementTenancy" : String,
			// "RamDiskId" : String,
			"SecurityGroups": ArrayOf(Schema{Type: TypeString}),
			// "SpotPrice" : String,
			"UserData": Schema{Type: TypeString},
		},
	}
}
