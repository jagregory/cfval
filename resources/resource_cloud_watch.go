package resources

import . "github.com/jagregory/cfval/schema"

var metricDimension = NestedResource{
	Description: "CloudWatch Alarm MetricDimension",
	Properties: Properties{
		"Name": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: true,
		},
	},
}

func Alarm() Resource {
	return Resource{
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
				Required: true,
				Type:     EnumValue{[]string{"GreaterThanOrEqualToThreshold", "GreaterThanThreshold", "LessThanThreshold", "LessThanOrEqualToThreshold"}},
			},

			"Dimensions": Schema{
				Type:  metricDimension,
				Array: true,
			},

			"EvaluationPeriods": Schema{
				Type:     ValueString,
				Required: true,
			},

			"InsufficientDataActions": Schema{
				Type:  ValueString,
				Array: true,
			},

			"MetricName": Schema{
				Type:     ValueString,
				Required: true,
			},

			"Namespace": Schema{
				Type:     ValueString,
				Required: true,
			},

			"OKActions": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Period": Schema{
				Type:     period,
				Required: true,
			},

			"Statistic": Schema{
				Type:     EnumValue{[]string{"SampleCount", "Average", "Sum", "Minimum", "Maximum"}},
				Required: true,
			},

			"Threshold": Schema{
				Type:     ValueString,
				Required: true,
			},

			"Unit": Schema{
				Type: EnumValue{[]string{"Seconds", "Microseconds", "Milliseconds", "Bytes", "Kilobytes", "Megabytes", "Gigabytes", "Terabytes", "Bits", "Kilobits", "Megabits", "Gigabits", "Terabits", "Percent", "Count", "Bytes/Second", "Kilobytes/Second", "Megabytes/Second", "Gigabytes/Second", "Terabytes/Second", "Bits/Second", "Kilobits/Second", "Megabits/Second", "Gigabits/Second", "Terabits/Second", "Count/Second", "None"}},
			},
		},
	}
}
