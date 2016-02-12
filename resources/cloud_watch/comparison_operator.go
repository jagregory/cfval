package cloud_watch

import . "github.com/jagregory/cfval/schema"

var comparisonOperator = EnumValue{
	Description: "Alarm Comparison Operator",

	Options: []string{"GreaterThanOrEqualToThreshold", "GreaterThanThreshold", "LessThanThreshold", "LessThanOrEqualToThreshold"},
}
