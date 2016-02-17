package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html
var Distribution = Resource{
	AwsType: "AWS::CloudFront::Distribution",

	Attributes: map[string]Schema{
		"DomainName": Schema{
			Type: ValueString,
		},
	},

	// Distribution ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DistributionConfig": Schema{
			Required: constraints.Always,
			Type:     distributionConfig,
		},
	},
}
