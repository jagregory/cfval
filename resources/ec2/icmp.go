package ec2

import . "github.com/jagregory/cfval/schema"
import "github.com/jagregory/cfval/constraints"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-icmp.html
var icmp = NestedResource{
	Description: "EC2 ICMP",

	Properties: Properties{
		"Code": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"Type": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
