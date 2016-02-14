package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-port-range.html
var portRange = NestedResource{
	Description: "EC2 PortRange",

	Properties: Properties{
		"From": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"To": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
