package resources

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var aliasTarget = NestedResource{
	Description: "Route53 RecordSet AliasTarget",
	Properties: map[string]Schema{
		"DNSName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"EvaluateTargetHealth": Schema{
			Type: ValueBool,
		},

		"HostedZoneId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

var continentCode = EnumValue{
	Description: "GeoLocation Continent Code",

	Options: []string{"AF", "AN", "AS", "EU", "OC", "NA", "SA"},
}

var countryCode = EnumValue{
	Description: "GeoLocation Country Code",

	Options: []string{"AO", "BF", "BI", "BJ", "BW", "CD", "CF", "CG", "CI", "CM", "CV", "DJ", "DZ", "EG", "ER", "ET", "GA", "GH", "GM", "GN", "GQ", "GW", "KE", "KM", "LR", "LS", "LY", "MA", "MG", "ML", "MR", "MU", "MW", "MZ", "NA", "NE", "NG", "RE", "RW", "SC", "SD", "SH", "SL", "SN", "SO", "SS", "ST", "SZ", "TD", "TG", "TN", "TZ", "UG", "YT", "ZA", "ZM", "ZW",
		"AQ", "GS", "TF",
		"AE", "AF", "AM", "AZ", "BD", "BH", "BN", "BT", "CC", "CN", "GE", "HK", "ID", "IL", "IN", "IO", "IQ", "IR", "JO", "JP", "KG", "KH", "KP", "KR", "KW", "KZ", "LA", "LB", "LK", "MM", "MN", "MO", "MV", "MY", "NP", "OM", "PH", "PK", "PS", "QA", "SA", "SG", "SY", "TH", "TJ", "TM", "TR", "TW", "UZ", "VN", "YE",
		"AD", "AL", "AT", "AX", "BA", "BE", "BG", "BY", "CH", "CY", "CZ", "DE", "DK", "EE", "ES", "FI", "FO", "FR", "GB", "GG", "GI", "GR", "HR", "HU", "IE", "IM", "IS", "IT", "JE", "LI", "LT", "LU", "LV", "MC", "MD", "ME", "MK", "MT", "NL", "NO", "PL", "PT", "RO", "RS", "RU", "SE", "SI", "SJ", "SK", "SM", "UA", "VA", "XK",
		"AG", "AI", "AW", "BB", "BL", "BM", "BQ", "BS", "BZ", "CA", "CR", "CU", "CW", "DM", "DO", "GD", "GL", "GP", "GT", "HN", "HT", "JM", "KN", "KY", "LC", "MF", "MQ", "MS", "MX", "NI", "PA", "PM", "PR", "SV", "SX", "TC", "TT", "US", "VC", "VG", "VI",
		"AS", "AU", "CK", "FJ", "FM", "GU", "KI", "MH", "MP", "NC", "NF", "NR", "NU", "NZ", "PF", "PG", "PN", "PW", "SB", "TK", "TL", "TO", "TV", "UM", "VU", "WF", "WS",
		"AR", "BO", "BR", "CL", "CO", "EC", "FK", "GF", "GY", "PE", "PY", "SR", "UY", "VE"},
}

var subdivisionCode = EnumValue{
	Description: "GeoLocation Subdivision Code",

	Options: []string{"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME", "MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM", "NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VA", "VT", "WA", "WI", "WV", "WY"},
}

var geoLocation = NestedResource{
	Description: "Route 53 Record Set GeoLocation",
	Properties: map[string]Schema{
		"ContinentCode": Schema{
			Type:     continentCode,
			Required: constraints.PropertyNotExists("CountryCode"),
			Conflicts: constraints.Any{
				constraints.PropertyExists("CountryCode"),
				constraints.PropertyExists("SubdivisionCode"),
			},
		},

		"CountryCode": Schema{
			Type:      countryCode,
			Required:  constraints.PropertyNotExists("ContinentCode"),
			Conflicts: constraints.PropertyExists("ContinentCode"),
		},

		"SubdivisionCode": Schema{
			Type:      subdivisionCode,
			Conflicts: constraints.PropertyExists("ContinentCode"),
			ValidateFunc: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
				if countryCode, found := self.Property("CountryCode"); found && countryCode != "US" {
					return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Can only be set when CountryCode is US", context)}
				}

				return reporting.ValidateOK, nil
			},
		},
	},
}

var recordSetType = EnumValue{
	Description: "RecordSet Type",

	Options: []string{"A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT"},
}

func RecordSet() Resource {
	return Resource{
		AwsType: "AWS::Route53::RecordSet",

		// Domain Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AliasTarget": Schema{
				Type: aliasTarget,
				Conflicts: constraints.Any{
					constraints.PropertyExists("ResourceRecords"),
					constraints.PropertyExists("TTL"),
				},
			},

			"Failover": Schema{
				Type: ValueString,
			},

			"GeoLocation": Schema{
				Type: geoLocation,
			},

			// "HealthCheckId":   Schema{Type: TypeString},

			"HostedZoneId": Schema{
				Type:      ValueString,
				Required:  constraints.PropertyNotExists("HostedZoneName"),
				Conflicts: constraints.PropertyExists("HostedZoneName"),
			},

			"HostedZoneName": Schema{
				Type:      ValueString,
				Required:  constraints.PropertyNotExists("HostedZoneId"),
				Conflicts: constraints.PropertyExists("HostedZoneId"),
			},

			"Name": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			// "Region":          Schema{Type: TypeString},

			"ResourceRecords": Schema{
				Array:     true,
				Type:      ValueString,
				Conflicts: constraints.PropertyExists("AliasTarget"),
				Required:  constraints.PropertyNotExists("AliasTarget"),
			},

			"SetIdentifier": Schema{
				Type: ValueString,
				Required: constraints.Any{
					constraints.PropertyExists("Weight"),
					constraints.PropertyExists("Latency"),
					constraints.PropertyExists("Failover"),
					constraints.PropertyExists("GeoLocation"),
				},
			},

			"TTL": Schema{
				Type:      ValueString,
				Conflicts: constraints.PropertyExists("AliasTarget"),
			},

			"Type": Schema{
				Type:     recordSetType,
				Required: constraints.Always,
			},

			"Weight": Schema{
				Type: ValueNumber,
			},
		},
	}
}
