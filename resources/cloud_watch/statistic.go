package cloud_watch

import . "github.com/jagregory/cfval/schema"

var statistic = EnumValue{
	Description: "Alarm Statistic",

	Options: []string{"SampleCount", "Average", "Sum", "Minimum", "Maximum"},
}
