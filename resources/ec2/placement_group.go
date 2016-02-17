package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
var PlacementGroup = Resource{
	AwsType: "AWS::EC2::PlacementGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Strategy": Schema{
			Type:         ValueString,
			Default:      "cluster",
			ValidateFunc: SingleValueValidate("cluster"),
		},
	},
}
