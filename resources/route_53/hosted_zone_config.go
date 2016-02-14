package route_53

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
var hostedZoneConfig = NestedResource{
	Description: "Route 53 HostedZoneConfig",

	Properties: Properties{
		"Comment": Schema{
			Type: ValueString,
		},
	},
}
