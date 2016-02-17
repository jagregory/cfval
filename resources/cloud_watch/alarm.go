package cloud_watch

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html
var Alarm = Resource{
	AwsType: "AWS::CloudWatch::Alarm",

	// AlarmName
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"ActionsEnabled": Schema{
			Type: ValueBool,
		},

		"AlarmActions": Schema{
			Array: true,
			Type:  ValueString,
		},

		"AlarmDescription": Schema{
			Type: ValueString,
		},

		"AlarmName": Schema{
			Type: ValueString,
		},

		"ComparisonOperator": Schema{
			Required: constraints.Always,
			Type:     comparisonOperator,
		},

		"Dimensions": Schema{
			Type:  metricDimension,
			Array: true,
		},

		"EvaluationPeriods": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"InsufficientDataActions": Schema{
			Type:  ValueString,
			Array: true,
		},

		"MetricName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Namespace": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OKActions": Schema{
			Type:  ValueString,
			Array: true,
		},

		"Period": Schema{
			Type:     Period,
			Required: constraints.Always,
		},

		"Statistic": Schema{
			Type:     statistic,
			Required: constraints.Always,
		},

		"Threshold": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Unit": Schema{
			Type: unit,
		},
	},
}
