package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
var healthCheck = NestedResource{
	Description: "ElasticLoadBalancing HealthCheck",
	Properties: Properties{
		"HealthyThreshold": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Interval": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Target": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		}, // TODO: Could be smarter about this restriction: "The protocol can be TCP, HTTP, HTTPS, or SSL. The range of valid ports is 1 through 65535."

		"Timeout": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		}, // TODO: Could be smarter about this restriction: "This value must be less than the value for Interval."

		"UnhealthyThreshold": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
