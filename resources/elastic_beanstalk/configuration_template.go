package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

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

			// "EnvironmentId": Schema{Type:TypeString},

			"OptionSettings": Schema{
				Type:  optionsSettings,
				Array: true,
			},

			"SolutionStackName": Schema{
				Type: ValueString,
			},

			// "SourceConfiguration": Schema{Type:...},
		},
	}
}
