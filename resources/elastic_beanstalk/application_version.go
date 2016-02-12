package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func ApplicationVersion() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ApplicationVersion",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"ApplicationName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Description": Schema{
				Type: ValueString,
			},

			"SourceBundle": Schema{
				Required: constraints.Always,
				Type: NestedResource{
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
				},
			},
		},
	}
}
