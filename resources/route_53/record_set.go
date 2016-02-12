package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

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
