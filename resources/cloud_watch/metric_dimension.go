package cloud_watch

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
var metricDimension = NestedResource{
	Description: "CloudWatch Alarm MetricDimension",
	Properties: Properties{
		"Name": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: StringLengthValidate(1, 255),
		},

		"Value": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: StringLengthValidate(1, 255),
		},
	},
}
