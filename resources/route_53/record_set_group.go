package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup.html
func RecordSetGroup() Resource {
	return Resource{
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
				Type:     recordSet,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
