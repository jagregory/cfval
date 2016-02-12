package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html
var sourceBundle = NestedResource{
	Description: "Elastic Beanstalk SourceBundle",
	Properties: Properties{
		"S3Bucket": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"S3Key": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
