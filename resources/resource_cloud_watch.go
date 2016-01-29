package resources

import . "github.com/jagregory/cfval/schema"

var metricDimension = Resource{
	AwsType: "CloudWatch Alarm MetricDimension",
	Properties: map[string]Schema{
		"Name":  Schema{Type: TypeString, Required: true},
		"Value": Schema{Type: TypeString, Required: true},
	},
}

func Alarm() Resource {
	return Resource{
		AwsType: "AWS::CloudWatch::Alarm",
		Properties: map[string]Schema{
			"ActionsEnabled":          Schema{Type: TypeBool},
			"AlarmActions":            ArrayOf(Schema{Type: TypeString}),
			"AlarmDescription":        Schema{Type: TypeString},
			"AlarmName":               Schema{Type: TypeString},
			"ComparisonOperator":      Required(EnumOf("GreaterThanOrEqualToThreshold", "GreaterThanThreshold", "LessThanThreshold", "LessThanOrEqualToThreshold")),
			"Dimensions":              ArrayOf(Schema{Type: metricDimension}),
			"EvaluationPeriods":       Required(Schema{Type: TypeString}),
			"InsufficientDataActions": ArrayOf(Schema{Type: TypeString}),
			"MetricName":              Schema{Type: TypeString, Required: true},
			"Namespace":               Schema{Type: TypeString, Required: true},
			"OKActions":               ArrayOf(Schema{Type: TypeString}),
			"Period":                  Required(period),
			"Statistic":               Required(EnumOf("SampleCount", "Average", "Sum", "Minimum", "Maximum")),
			"Threshold":               Schema{Type: TypeString, Required: true},
			"Unit":                    EnumOf("Seconds", "Microseconds", "Milliseconds", "Bytes", "Kilobytes", "Megabytes", "Gigabytes", "Terabytes", "Bits", "Kilobits", "Megabits", "Gigabits", "Terabits", "Percent", "Count", "Bytes/Second", "Kilobytes/Second", "Megabytes/Second", "Gigabytes/Second", "Terabytes/Second", "Bits/Second", "Kilobits/Second", "Megabits/Second", "Gigabits/Second", "Terabits/Second", "Count/Second", "None"),
		},
	}
}
