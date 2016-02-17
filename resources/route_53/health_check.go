package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
var HealthCheck = Resource{
	AwsType: "AWS::Route53::HealthCheck",

	Properties: Properties{
		"HealthCheckConfig": Schema{
			Type:     healthCheckConfig,
			Required: constraints.Always,
		},

		"HealthCheckTags": Schema{
			Type:  healthCheckTag,
			Array: true,
		},
	},
}
