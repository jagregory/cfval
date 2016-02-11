package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
var ssmAssociation = NestedResource{
	Description: "EC2 Instance SsmAssociation",
	Properties: Properties{
		"AssociationParameters": Schema{
			Type:  associationParameter,
			Array: true,
		},

		"DocumentName": Schema{
			Type:     ValueString,
			Required: Always,
		},
	},
}
