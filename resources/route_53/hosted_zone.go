package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
var HostedZone = Resource{
	AwsType: "AWS::Route53::HostedZone",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"HostedZoneConfig": Schema{
			Type: hostedZoneConfig,
		},

		"HostedZoneTags": Schema{
			Type:  hostedZoneTag,
			Array: true,
		},

		"Name": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"VPCs": Schema{
			Type:  hostedZoneVPC,
			Array: true,
		},
	},
}
