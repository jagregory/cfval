package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html
var stepAdjustment = NestedResource{
	Description: "Auto Scaling ScalingPolicy StepAdjustments",

	Properties: Properties{
		"MetricIntervalLowerBound": Schema{
			Type: ValueNumber,
		},

		"MetricIntervalUpperBound": Schema{
			Type: ValueNumber,
		},

		"ScalingAdjustment": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
