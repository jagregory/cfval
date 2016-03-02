package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
var ScalingPolicy = Resource{
	AwsType: "AWS::AutoScaling::ScalingPolicy",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AdjustmentType": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"AutoScalingGroupName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Cooldown": Schema{
			Type:      ValueString,
			Conflicts: constraints.PropertyIs("PolicyType", "StepScaling"),
		},

		"EstimatedInstanceWarmup": Schema{
			Type:      ValueNumber,
			Conflicts: constraints.PropertyIs("PolicyType", "SimpleScaling"),
		},

		"MetricAggregationType": Schema{
			Type:      metricAggregationType,
			Default:   "Average",
			Conflicts: constraints.PropertyIs("PolicyType", "SimpleScaling"),
		},

		// TODO: This property replaces the MinAdjustmentStep property
		"MinAdjustmentMagnitude": Schema{
			Type: ValueNumber,
		},

		"PolicyType": Schema{
			Type:    policyType,
			Default: "SimpleScaling",
		},

		"ScalingAdjustment": Schema{
			Type:      ValueNumber,
			Required:  constraints.PropertyIs("PolicyType", "SimpleScaling"),
			Conflicts: constraints.Not(constraints.PropertyIs("PolicyType", "SimpleScaling")),
		},

		"StepAdjustments": Schema{
			Type:      Multiple(stepAdjustment),
			Required:  constraints.PropertyIs("PolicyType", "StepScaling"),
			Conflicts: constraints.PropertyNot("PolicyType", "StepScaling"),
		},
	},
}
