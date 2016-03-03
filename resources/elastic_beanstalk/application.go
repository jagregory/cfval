package elastic_beanstalk

import (
	"github.com/jagregory/cfval/deprecations"
	. "github.com/jagregory/cfval/schema"
)

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

		"ApplicationVersions": Schema{
			Deprecated: deprecations.Deprecated("Application Versions should be created by using the separate AWS::ElasticBeanstalk::ApplicationVersion resource, and linked using the ApplicationName property."),
		},

		"ConfigurationTemplates": Schema{
			Deprecated: deprecations.Deprecated("Configuration Templates should be created by using the separate AWS::ElasticBeanstalk::ConfigurationTemplate resource, and linked using the ApplicationName property."),
		},

		"Description": Schema{
			Type: ValueString,
		},
	},
}
