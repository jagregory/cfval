package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
var SubnetNetworkACLAssociation = Resource{
	AwsType: "AWS::EC2::SubnetNetworkAclAssociation",

	Attributes: map[string]Schema{
		"AssociationId": Schema{
			Type: SubnetID,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"SubnetId": Schema{
			Type:     SubnetID,
			Required: constraints.Always,
		},

		"NetworkAclId": Schema{
			Type:     NetworkAclID,
			Required: constraints.Always,
		},
	},
}
