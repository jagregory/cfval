package resources

import . "github.com/jagregory/cfval/schema"

var metricDimension = Resource{
	AwsType: "CloudWatch Alarm MetricDimension",
	Properties: map[string]Schema{
		"Name": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Value": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

func Alarm() Resource {
	return Resource{
		AwsType: "AWS::CloudWatch::Alarm",
		Properties: map[string]Schema{
			"ActionsEnabled": Schema{
				Type: TypeBool,
			},

			"AlarmActions": Schema{
				Array: true,
				Type:  TypeString,
			},

			"AlarmDescription": Schema{
				Type: TypeString,
			},

			"AlarmName": Schema{
				Type: TypeString,
			},

			"ComparisonOperator": Schema{
				Required:     true,
				Type:         TypeString,
				ValidateFunc: EnumValidate("GreaterThanOrEqualToThreshold", "GreaterThanThreshold", "LessThanThreshold", "LessThanOrEqualToThreshold"),
			},

			"Dimensions": Schema{
				Type:  metricDimension,
				Array: true,
			},

			"EvaluationPeriods": Schema{
				Type:     TypeString,
				Required: true,
			},

			"InsufficientDataActions": Schema{
				Type:  TypeString,
				Array: true,
			},

			"MetricName": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Namespace": Schema{
				Type:     TypeString,
				Required: true,
			},

			"OKActions": Schema{
				Type:  TypeString,
				Array: true,
			},

			"Period": Schema{
				Type:         TypeString,
				Required:     true,
				ValidateFunc: period,
			},

			"Statistic": Schema{
				Type:         TypeString,
				Required:     true,
				ValidateFunc: EnumValidate("SampleCount", "Average", "Sum", "Minimum", "Maximum"),
			},

			"Threshold": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Unit": Schema{
				Type:         TypeString,
				ValidateFunc: EnumValidate("Seconds", "Microseconds", "Milliseconds", "Bytes", "Kilobytes", "Megabytes", "Gigabytes", "Terabytes", "Bits", "Kilobits", "Megabits", "Gigabits", "Terabits", "Percent", "Count", "Bytes/Second", "Kilobytes/Second", "Megabytes/Second", "Gigabytes/Second", "Terabytes/Second", "Bits/Second", "Kilobits/Second", "Megabits/Second", "Gigabits/Second", "Terabits/Second", "Count/Second", "None"),
			},
		},
	}
}
