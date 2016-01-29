package resources

import . "github.com/jagregory/cfval/schema"

var metricsCollection = Schema{
	Type: Resource{
		AwsType:    "Auto Scaling MetricsCollection",
		Properties: map[string]Schema{
		// TODO
		},
	},
}
