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
			Type:     Multiple(ValueString),
			Required: constraints.Always,
		},

		"LogGroupName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"MetricTransformations": Schema{
			Type:     Multiple(metricTransformation),
			Required: constraints.Always,
		},
	},
}
