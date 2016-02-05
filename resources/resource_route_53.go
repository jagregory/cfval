package resources

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

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

var geoLocation = Resource{
	AwsType: "Route 53 Record Set GeoLocation",
	Properties: map[string]Schema{
		"ContinentCode": Schema{
			Type:           TypeString,
			RequiredUnless: []string{"CountryCode"},
			Conflicts:      []string{"CountryCode", "SubdivisionCode"},
			ValidateFunc:   EnumValidate("AF", "AN", "AS", "EU", "OC", "NA", "SA"),
		},

		"CountryCode": Schema{
			Type:           TypeString,
			RequiredUnless: []string{"ContinentCode"},
			Conflicts:      []string{"ContinentCode"},
			ValidateFunc: EnumValidate("AO", "BF", "BI", "BJ", "BW", "CD", "CF", "CG", "CI", "CM", "CV", "DJ", "DZ", "EG", "ER", "ET", "GA", "GH", "GM", "GN", "GQ", "GW", "KE", "KM", "LR", "LS", "LY", "MA", "MG", "ML", "MR", "MU", "MW", "MZ", "NA", "NE", "NG", "RE", "RW", "SC", "SD", "SH", "SL", "SN", "SO", "SS", "ST", "SZ", "TD", "TG", "TN", "TZ", "UG", "YT", "ZA", "ZM", "ZW",
				"AQ", "GS", "TF",
				"AE", "AF", "AM", "AZ", "BD", "BH", "BN", "BT", "CC", "CN", "GE", "HK", "ID", "IL", "IN", "IO", "IQ", "IR", "JO", "JP", "KG", "KH", "KP", "KR", "KW", "KZ", "LA", "LB", "LK", "MM", "MN", "MO", "MV", "MY", "NP", "OM", "PH", "PK", "PS", "QA", "SA", "SG", "SY", "TH", "TJ", "TM", "TR", "TW", "UZ", "VN", "YE",
				"AD", "AL", "AT", "AX", "BA", "BE", "BG", "BY", "CH", "CY", "CZ", "DE", "DK", "EE", "ES", "FI", "FO", "FR", "GB", "GG", "GI", "GR", "HR", "HU", "IE", "IM", "IS", "IT", "JE", "LI", "LT", "LU", "LV", "MC", "MD", "ME", "MK", "MT", "NL", "NO", "PL", "PT", "RO", "RS", "RU", "SE", "SI", "SJ", "SK", "SM", "UA", "VA", "XK",
				"AG", "AI", "AW", "BB", "BL", "BM", "BQ", "BS", "BZ", "CA", "CR", "CU", "CW", "DM", "DO", "GD", "GL", "GP", "GT", "HN", "HT", "JM", "KN", "KY", "LC", "MF", "MQ", "MS", "MX", "NI", "PA", "PM", "PR", "SV", "SX", "TC", "TT", "US", "VC", "VG", "VI",
				"AS", "AU", "CK", "FJ", "FM", "GU", "KI", "MH", "MP", "NC", "NF", "NR", "NU", "NZ", "PF", "PG", "PN", "PW", "SB", "TK", "TL", "TO", "TV", "UM", "VU", "WF", "WS",
				"AR", "BO", "BR", "CL", "CO", "EC", "FK", "GF", "GY", "PE", "PY", "SR", "UY", "VE"),
		},

		"SubdivisionCode": Schema{
			Type:      TypeString,
			Conflicts: []string{"ContinentCode"},
			ValidateFunc: func(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
				if !tr.HasProperty("CountryCode", "US") {
					return false, []reporting.Failure{reporting.NewFailure("Can only be set when CountryCode is US", context)}
				}

				return EnumValidate("AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME", "MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM", "NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VA", "VT", "WA", "WI", "WV", "WY")(value, tr, context)
			},
		},
	},
}

func RecordSet() Resource {
	return Resource{
		AwsType: "AWS::Route53::RecordSet",

		// Domain Name
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"AliasTarget": Schema{
				Type:      aliasTarget,
				Conflicts: []string{"ResourceRecords", "TTL"},
			},

			"Failover": Schema{
				Type: TypeString,
			},

			"GeoLocation": Schema{
				Type: geoLocation,
			},

			// "HealthCheckId":   Schema{Type: TypeString},

			"HostedZoneId": Schema{
				Type:           TypeString,
				RequiredUnless: []string{"HostedZoneName"},
				Conflicts:      []string{"HostedZoneName"},
			},

			"HostedZoneName": Schema{
				Type:           TypeString,
				RequiredUnless: []string{"HostedZoneId"},
				Conflicts:      []string{"HostedZoneId"},
			},

			"Name": Schema{
				Type:     TypeString,
				Required: true,
			},

			// "Region":          Schema{Type: TypeString},

			"ResourceRecords": Schema{
				Array:          true,
				Type:           TypeString,
				Conflicts:      []string{"AliasTarget"},
				RequiredUnless: []string{"AliasTarget"},
			},

			"SetIdentifier": Schema{
				Type:       TypeString,
				RequiredIf: []string{"Weight", "Latency", "Failover", "GeoLocation"},
			},

			"TTL": Schema{
				Type:      TypeString,
				Conflicts: []string{"AliasTarget"},
			},

			"Type": Schema{
				Type:         TypeString,
				Required:     true,
				ValidateFunc: EnumValidate("A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT"),
			},

			"Weight": Schema{
				Type: TypeInteger,
			},
		},
	}
}
