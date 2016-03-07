package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

var recordSetGroupRecordSet = NestedResource{
	Description: "RecordSetGroup RecordSet",
	Properties:  recordSetProperties(false),
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup.html
var RecordSetGroup = Resource{
	AwsType: "AWS::Route53::RecordSetGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Comment": Schema{
			Type: ValueString,
		},

		"HostedZoneId": Schema{
			Type:     HostedZoneID,
			Required: constraints.PropertyNotExists("HostedZoneName"),
		},

		"HostedZoneName": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("HostedZoneId"),
		},

		"RecordSets": Schema{
			Type:     Multiple(recordSetGroupRecordSet),
			Required: constraints.Always,
		},
	},
}
