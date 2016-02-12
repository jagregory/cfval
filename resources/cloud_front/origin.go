package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
var origin = NestedResource{
	Description: "CloudFront DistributionConfig Origin",
	Properties: Properties{
		"CustomOriginConfig": Schema{
			Type:      customOriginConfig,
			Conflicts: constraints.PropertyExists("S3OriginConfig"),
			Required:  constraints.PropertyNotExists("S3OriginConfig"),
		},

		"DomainName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Id": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OriginPath": Schema{
			Type:         ValueString,
			ValidateFunc: RegexpValidate(`^\/.*?[^\/]$`, "The value must start with a slash mark (/) and cannot end with a slash mark."),
		},

		"S3OriginConfig": Schema{
			Type:      s3OriginConfig,
			Conflicts: constraints.PropertyExists("CustomOriginConfig"),
			Required:  constraints.PropertyNotExists("CustomOriginConfig"),
		},
	},
}
