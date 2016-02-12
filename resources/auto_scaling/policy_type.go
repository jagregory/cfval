package auto_scaling

import . "github.com/jagregory/cfval/schema"

var policyType = EnumValue{
	Description: "ScalingPolicy PolicyType",

	Options: []string{"SimpleScaling", "StepScaling"},
}
