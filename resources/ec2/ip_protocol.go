package ec2

import . "github.com/jagregory/cfval/schema"

var ipProtocol = EnumValue{
	Description: "SecurityGroupIngress IpProtocol",

	Options: []string{"tcp", "udp", "icmp", "-1"},
}
