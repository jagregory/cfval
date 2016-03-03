package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var recordSetProperties = Properties{
	"AliasTarget": Schema{
		Type: aliasTarget,
		Conflicts: constraints.Any{
			constraints.PropertyExists("ResourceRecords"),
			constraints.PropertyExists("TTL"),
		},
	},

	"Failover": Schema{
		Type: EnumValue{
			Description: "Failover",
			Options:     []string{"PRIMARY", "SECONDARY"},
		},
	},

	"GeoLocation": Schema{
		Type: geoLocation,
	},

	"HealthCheckId": Schema{
		Type: ValueString,
	},

	"HostedZoneId": Schema{
		Type:      HostedZoneID,
		Conflicts: constraints.PropertyExists("HostedZoneName"),
	},

	"HostedZoneName": Schema{
		Type:      ValueString,
		Conflicts: constraints.PropertyExists("HostedZoneId"),
	},

	"Name": Schema{
		Type:     ValueString,
		Required: constraints.Always,
	},

	// TODO: Region validation: http://docs.aws.amazon.com/general/latest/gr/rande.html
	"Region": Schema{
		Type: ValueString,
	},

	"ResourceRecords": Schema{
		Type:      Multiple(ValueString),
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
}

var recordSet = NestedResource{
	Description: "RecordSetGroup RecordSet",
	Properties:  recordSetProperties,
}

var RecordSet = Resource{
	AwsType: "AWS::Route53::RecordSet",

	// Domain Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: recordSetProperties,
}
