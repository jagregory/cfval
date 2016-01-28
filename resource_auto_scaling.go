package main

func autoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]Schema{
			"AvailabilityZones":          ArrayOf(AvailabilityZone),
			"Cooldown":                   Schema{Type: TypeString},
			"DesiredCapacity":            Schema{Type: TypeString},
			"HealthCheckGracePeriod":     Schema{Type: TypeInteger},
			"HealthCheckType":            EnumSchema("EC2", "ELB"),
			"InstanceId":                 Schema{Type: TypeString},
			"LaunchConfigurationName":    Schema{Type: TypeString},
			"LoadBalancerNames":          ArrayOf(Schema{Type: TypeString}),
			"MaxSize":                    Schema{Type: TypeString},
			"MetricsCollection":          ArrayOf(MetricsCollection),
			"MinSize":                    Schema{Type: TypeString},
			"NotificationConfigurations": ArrayOf(NotificationConfiguration),
			"PlacementGroup":             Schema{Type: TypeString},
			"Tags":                       ArrayOf(AutoScalingTag),
			"TerminationPolicies":        ArrayOf(Schema{Type: TypeString}),
			"VPCZoneIdentifier":          ArrayOf(Schema{Type: TypeString}),
		},
	}
}

func launchConfiguration() Resource {
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
