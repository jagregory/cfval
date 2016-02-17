package elastic_beanstalk

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
var Application = Resource{
	AwsType: "AWS::ElasticBeanstalk::Application",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"ApplicationName": Schema{
			Type: ValueString,
		},

		"Description": Schema{
			Type: ValueString,
		},
	},
}
