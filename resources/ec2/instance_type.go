package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html
var ec2InstanceType = EnumValue{
	Description: "EC2 Instance Type",
	Options: []string{
		// Current Generation Instances
		// General purpose
		"t2.nano", "t2.micro", "t2.small", "t2.medium", "t2.large", "m4.large", "m4.xlarge", "m4.2xlarge", "m4.4xlarge", "m4.10xlarge", "m3.medium", "m3.large", "m3.xlarge", "m3.2xlarge",
		// Compute optimized
		"c4.large", "c4.xlarge", "c4.2xlarge", "c4.4xlarge", "c4.8xlarge", "c3.large", "c3.xlarge", "c3.2xlarge", "c3.4xlarge", "c3.8xlarge",
		// Memory optimized
		"r3.large", "r3.xlarge", "r3.2xlarge", "r3.4xlarge", "r3.8xlarge",
		// Storage optimized
		"i2.xlarge", "i2.2xlarge", "i2.4xlarge", "i2.8xlarge", "d2.xlarge", "d2.2xlarge", "d2.4xlarge", "d2.8xlarge",
		// GPU instances
		"g2.2xlarge", "g2.8xlarge",

		// Previous Generation Instances
		// General purpose
		"m1.small", "m1.medium", "m1.large", "m1.xlarge",
		// Compute optimized
		"c1.medium", "c1.xlarge", "cc2.8xlarge",
		// Memory optimized
		"m2.xlarge", "m2.2xlarge", "m2.4xlarge", "cr1.8xlarge",
		// Storage optimized
		"hi1.4xlarge", "hs1.8xlarge",
		// GPU instances
		"cg1.4xlarge",
		// Micro instances
		"t1.micro",
	},
}
