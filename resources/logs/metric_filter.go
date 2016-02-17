package logs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
var MetricFilter = Resource{
	AwsType: "AWS::Logs::MetricFilter",

	Properties: Properties{
		"FilterPattern": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.Always,
		},

		"LogGroupName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"MetricTransformations": Schema{
			Type:     metricTransformation,
			Array:    true,
			Required: constraints.Always,
		},
	},
}
