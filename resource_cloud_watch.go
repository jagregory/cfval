package main

var Period = Schema{
	Type: TypeString,
	ValidateFunc: func(interface{}, Template, []string) (bool, []Failure) {
		// TODO: The time over which the specified statistic is applied. You must specify a time in seconds that is also a multiple of 60.
		return true, nil
	},
}

var MetricDimension = Schema{
	Type: Resource{
		AwsType: "CloudWatch Alarm MetricDimension",
		Properties: map[string]Schema{
			"Name":  Schema{Type: TypeString, Required: true},
			"Value": Schema{Type: TypeString, Required: true},
		},
	},
}

func alarm() Resource {
	return Resource{
		AwsType: "AWS::CloudWatch::Alarm",
		Properties: map[string]Schema{
			"ActionsEnabled":          Schema{Type: TypeBool},
			"AlarmActions":            ArrayOf(Schema{Type: TypeString}),
			"AlarmDescription":        Schema{Type: TypeString},
			"AlarmName":               Schema{Type: TypeString},
			"ComparisonOperator":      Required(EnumSchema("GreaterThanOrEqualToThreshold", "GreaterThanThreshold", "LessThanThreshold", "LessThanOrEqualToThreshold")),
			"Dimensions":              ArrayOf(MetricDimension),
			"EvaluationPeriods":       Required(Schema{Type: TypeString}),
			"InsufficientDataActions": ArrayOf(Schema{Type: TypeString}),
			"MetricName":              Schema{Type: TypeString, Required: true},
			"Namespace":               Schema{Type: TypeString, Required: true},
			"OKActions":               ArrayOf(Schema{Type: TypeString}),
			"Period":                  Required(Period),
			"Statistic":               Required(EnumSchema("SampleCount", "Average", "Sum", "Minimum", "Maximum")),
			"Threshold":               Schema{Type: TypeString, Required: true},
			"Unit":                    EnumSchema("Seconds", "Microseconds", "Milliseconds", "Bytes", "Kilobytes", "Megabytes", "Gigabytes", "Terabytes", "Bits", "Kilobits", "Megabits", "Gigabits", "Terabits", "Percent", "Count", "Bytes/Second", "Kilobytes/Second", "Megabytes/Second", "Gigabytes/Second", "Terabytes/Second", "Bits/Second", "Kilobits/Second", "Megabits/Second", "Gigabits/Second", "Terabits/Second", "Count/Second", "None"),
		},
	}
}
