package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html
var spotFleetRequestConfigData = NestedResource{
	Description: "EC2 SpotFleet SpotFleetRequestConfigData",

	Properties: Properties{
		"AllocationStrategy": Schema{
			Type: EnumValue{
				Description: "AllocationStrategy",
				Options:     []string{"lowestPrice", "diversified"},
			},
		},

		"ExcessCapacityTerminationPolicy": Schema{
			Type: EnumValue{
				Description: "ExcessCapacityTerminationPolicy",
				Options:     []string{"noTermination", "default"},
			},
		},

		"IamFleetRole": Schema{
			Type:     ValueString, // ARN
			Required: constraints.Always,
		},

		"LaunchSpecifications": Schema{
			Type:     spotFleetRequestConfigDataLaunchSpecification,
			Array:    true,
			Required: constraints.Always,
		},

		"SpotPrice": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TargetCapacity": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"TerminateInstancesWithExpiration": Schema{
			Type: ValueBool,
		},

		"ValidFrom": Schema{
			Type: Timestamp,
		},

		"ValidUntil": Schema{
			Type: Timestamp,
		},
	},
}
