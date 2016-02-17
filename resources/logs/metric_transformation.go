package logs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
var metricTransformation = NestedResource{
	Description: "Logs MetricFilter MetricTransformation",

	Properties: Properties{
		"MetricName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"MetricNamespace": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"MetricValue": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
