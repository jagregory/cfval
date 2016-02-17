package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html
var DHCPOptions = Resource{
	AwsType: "AWS::EC2::DHCPOptions",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DomainName": Schema{
			Type: ValueString,
		},

		"DomainNameServers": Schema{
			Type:  IPAddress,
			Array: true,
			Required: constraints.All{
				constraints.PropertyNotExists("NetbiosNameServers"),
				constraints.PropertyNotExists("NtpServers"),
			},
		},

		"NetbiosNameServers": Schema{
			Type:  IPAddress,
			Array: true,
			Required: constraints.All{
				constraints.PropertyNotExists("DomainNameServers"),
				constraints.PropertyNotExists("NtpServers"),
			},
		},

		"NetbiosNodeType": Schema{
			Type:     ValueNumber,
			Required: constraints.PropertyExists("NetBiosNameServers"),
		},

		"NtpServers": Schema{
			Type:  IPAddress,
			Array: true,
			Required: constraints.All{
				constraints.PropertyNotExists("DomainNameServers"),
				constraints.PropertyNotExists("NetbiosNameServers"),
			},
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},
	},
}
