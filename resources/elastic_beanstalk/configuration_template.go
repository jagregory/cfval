package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html
func ConfigurationTemplate() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ConfigurationTemplate",

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

			"EnvironmentId": Schema{
				Type: ValueString,
				Required: constraints.All{
					constraints.PropertyNotExists("SolutionStackName"),
					constraints.PropertyNotExists("SourceConfiguration"),
				},
			},

			"OptionSettings": Schema{
				Type:  optionsSettings,
				Array: true,
			},

			"SolutionStackName": Schema{
				Type: ValueString,
				Required: constraints.All{
					constraints.PropertyNotExists("EnvironmentId"),
					constraints.PropertyNotExists("SourceConfiguration"),
				},
			},

			"SourceConfiguration": Schema{
				Type: sourceConfiguration,
				Required: constraints.All{
					constraints.PropertyNotExists("EnvironmentId"),
					constraints.PropertyNotExists("SolutionStackName"),
				},
			},
		},
	}
}
