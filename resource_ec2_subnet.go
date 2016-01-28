package main

func subnet() Resource {
	return Resource{
		Properties: map[string]Schema{
			"AvailabilityZone": AvailabilityZone,

			"CidrBlock": Required(Cidr),

			"MapPublicIpOnLaunch": Schema{
				Type: TypeBool,
			},

			"Tags": ArrayOf(ResourceTag),

			"VpcId": Required(VpcId),
		},
	}
}
