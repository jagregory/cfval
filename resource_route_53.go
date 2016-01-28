package main

var AliasTarget = Schema{
	Type: Resource{
		AwsType: "Route53 RecordSet AliasTarget",
		Properties: map[string]Schema{
			"DNSName":              Schema{Type: TypeString, Required: true},
			"EvaluateTargetHealth": Schema{Type: TypeBool},
			"HostedZoneId":         Schema{Type: TypeString, Required: true},
		},
	},
}

func recordSet() Resource {
	return Resource{
		AwsType: "AWS::Route53::RecordSet",
		Properties: map[string]Schema{
			"AliasTarget": AliasTarget,
			// "Failover": Schema{Type: TypeString},
			// "GeoLocation":     GeoLocation,
			// "HealthCheckId":   Schema{Type: TypeString},
			// "HostedZoneId":    Schema{Type: TypeString},
			"HostedZoneName": Schema{Type: TypeString},
			"Name":           Schema{Type: TypeString, Required: true},
			// "Region":          Schema{Type: TypeString},
			"ResourceRecords": ArrayOf(Schema{Type: TypeString}),
			// "SetIdentifier":   Schema{Type: TypeString},
			"TTL":  Schema{Type: TypeString},
			"Type": Required(EnumSchema("A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT")),
			// "Weight":          Schema{Type: TypeInteger},
		},
	}
}
