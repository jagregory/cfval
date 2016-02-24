package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
var metricsCollection = NestedResource{
	Description: "Auto Scaling MetricsCollection",
	Properties: Properties{
		"Granularity": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Metrics": Schema{
			Type: Multiple(ValueString),
		},
	},
}
