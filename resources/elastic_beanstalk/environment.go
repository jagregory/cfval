package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func Environment() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::Environment",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"ApplicationName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"CNAMEPrefix": Schema{
				Type: ValueString,
			},

			"Description": Schema{
				Type: ValueString,
			},

			"EnvironmentName": Schema{
				Type: ValueString,
			},

			"OptionSettings": Schema{
				Type:  optionsSettings,
				Array: true,
			},

			"SolutionStackName": Schema{
				Type: ValueString,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"TemplateName": Schema{
				Type: ValueString,
			},

			// "Tier": Schema{...}

			"VersionLabel": Schema{
				Type: ValueString,
			},
		},
	}
}
