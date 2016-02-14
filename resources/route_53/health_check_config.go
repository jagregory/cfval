package route_53

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html
var healthCheckConfig = NestedResource{
	Description: "Route 53 HealthCheckConfig",

	Properties: Properties{
		"FailureThreshold": Schema{
			Type: ValueNumber,
		},

		"FullyQualifiedDomainName": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("IPAddress"),
		},

		"IPAddress": Schema{
			Type:     IPAddress,
			Required: constraints.PropertyNotExists("FullyQualifiedDomainName"),
		},

		"Port": Schema{
			Type:     ValueNumber,
			Required: constraints.PropertyIs("Type", "TCP"),
		},

		"RequestInterval": Schema{
			Type: ValueNumber,
		},

		"ResourcePath": Schema{
			Type: ValueString,
		},

		"SearchString": Schema{
			Type: ValueString,
		},

		"Type": Schema{
			Type: EnumValue{
				Description: "Type",
				Options:     []string{"HTTP", "HTTPS", "HTTP_STR_MATCH", "HTTPS_STR_MATCH", "TCP"},
			},
		},
	},
}
