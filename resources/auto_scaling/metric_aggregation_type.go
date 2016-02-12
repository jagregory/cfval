package auto_scaling

import . "github.com/jagregory/cfval/schema"

var metricAggregationType = EnumValue{
	Description: "ScalingPolicy MetricAggregationType",

	Options: []string{"Minimum", "Maximum", "Average"},
}
