package main

var AutoScalingTag = Schema{
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
