package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
var SpotFleet = Resource{
	AwsType: "AWS::EC2::SpotFleet",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"SpotFleetRequestConfigData": Schema{
			Type:     spotFleetRequestConfigData,
			Required: constraints.Always,
		},
	},
}
