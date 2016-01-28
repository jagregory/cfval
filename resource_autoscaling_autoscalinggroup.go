package main

func autoScalingGroup() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]Schema{
			"AvailabilityZones": ArrayOf(AvailabilityZone),

			"Cooldown": Schema{
				Type: TypeString,
			},

			"DesiredCapacity": Schema{
				Type: TypeString,
			},

			"HealthCheckGracePeriod": Schema{
				Type: TypeInteger,
			},

			"Tags": ArrayOf(ResourceTag),
		},
	}
}
