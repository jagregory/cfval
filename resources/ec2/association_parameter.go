package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
var associationParameter = NestedResource{
	Description: "EC2 Instance SsmAssociations AssociationParameter",
	Properties: Properties{
		"Key": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Array:    true,
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
