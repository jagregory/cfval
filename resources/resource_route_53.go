package resources

import . "github.com/jagregory/cfval/schema"

var aliasTarget = Resource{
	AwsType: "Route53 RecordSet AliasTarget",
	Properties: map[string]Schema{
		"DNSName": Schema{
			Type:     TypeString,
			Required: true,
		},

		"EvaluateTargetHealth": Schema{
			Type: TypeBool,
		},

		"HostedZoneId": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

func RecordSet() Resource {
	return Resource{
		AwsType: "AWS::Route53::RecordSet",
		Properties: map[string]Schema{
			"AliasTarget": Schema{
				Type: aliasTarget,
			},

			// "Failover": Schema{Type: TypeString},
			// "GeoLocation":     GeoLocation,
			// "HealthCheckId":   Schema{Type: TypeString},
			// "HostedZoneId":    Schema{Type: TypeString},

			"HostedZoneName": Schema{
				Type: TypeString,
			},

			"Name": Schema{
				Type:     TypeString,
				Required: true,
			},

			// "Region":          Schema{Type: TypeString},

			"ResourceRecords": Schema{
				Type:     TypeString,
				Required: true,
			},

			// "SetIdentifier":   Schema{Type: TypeString},

			"TTL": Schema{
				Type: TypeString,
			},

			"Type": Required(EnumOf("A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT")),

			// "Weight":          Schema{Type: TypeInteger},
		},
	}
}
