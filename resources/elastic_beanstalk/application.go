package elastic_beanstalk

import . "github.com/jagregory/cfval/schema"

func Application() Resource {
	return Resource{
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
}
