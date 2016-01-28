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
			"ImageId":        Schema{Type: TypeString},
			"InstanceType":   Schema{Type: TypeString},
			"KeyName":        Schema{Type: TypeString},
			"SecurityGroups": ArrayOf(Schema{Type: TypeString}),
		},
	}
}
