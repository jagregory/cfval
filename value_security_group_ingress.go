package main

func securityGroupIngress() Resource {
	return Resource{
		AwsType: "AWS::EC2::SecurityGroupIngress",
		Properties: map[string]Schema{
			"CidrIp":                     Cidr,
			"FromPort":                   Schema{Type: TypeInteger, Required: true},
			"IpProtocol":                 Required(EnumSchema("tcp", "udp", "icmp", "-1")),
			"SourceSecurityGroupId":      Schema{Type: TypeString},
			"SourceSecurityGroupName":    Schema{Type: TypeString},
			"SourceSecurityGroupOwnerId": Schema{Type: TypeString},
			"ToPort":                     Schema{Type: TypeInteger, Required: true},
		},
	}
}

// TODO: this is a bit weird
var SecurityGroupIngress = Schema{
	Type: securityGroupIngress(),
}
