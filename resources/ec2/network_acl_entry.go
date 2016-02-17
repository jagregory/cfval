package ec2

import . "github.com/jagregory/cfval/schema"
import "github.com/jagregory/cfval/constraints"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html
var NetworkACLEntry = Resource{
	AwsType: "AWS::EC2::NetworkAclEntry",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CidrBlock": Schema{
			Type:     CIDR,
			Required: constraints.Always,
		},

		"Egress": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"Icmp": Schema{
			Type:     icmp,
			Required: constraints.PropertyIs("Protocol", 1), // ICMP
		},

		"NetworkAclId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PortRange": Schema{
			Type: portRange,
			Required: constraints.Any{
				constraints.PropertyIs("Protocol", 6),  // TCP
				constraints.PropertyIs("Protocol", 17), // UDP
			},
		},

		"Protocol": Schema{
			Type:         ValueNumber,
			Required:     constraints.Always,
			ValidateFunc: IntegerRangeValidate(-1, 255),
		},

		"RuleAction": Schema{
			Type: EnumValue{
				Description: "RuleAction",
				Options:     []string{"allow", "deny"},
			},
			Required: constraints.Always,
		},

		"RuleNumber": Schema{
			Type:         ValueNumber,
			Required:     constraints.Always,
			ValidateFunc: IntegerRangeValidate(1, 32766),
		},
	},
}
