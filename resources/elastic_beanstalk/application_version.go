package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
var ApplicationVersion = Resource{
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
			Type:     sourceBundle,
		},
	},
}
