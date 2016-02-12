package cloud_watch

import . "github.com/jagregory/cfval/schema"

var unit = EnumValue{
	Description: "Alarm Unit",

	Options: []string{
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Count",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
		"None",
	},
}
