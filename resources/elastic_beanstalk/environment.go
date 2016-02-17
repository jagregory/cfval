package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html
var Environment = Resource{
	AwsType: "AWS::ElasticBeanstalk::Environment",

	Attributes: map[string]Schema{
		"EndpointURL": Schema{
			Type: ValueString,
		},
	},

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
			Type:     ValueString,
			Required: constraints.PropertyNotExists("TemplateName"),
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},

		"TemplateName": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("SolutionStackName"),
		},

		"Tier": Schema{
			Type: tier,
		},

		"VersionLabel": Schema{
			Type: ValueString,
		},
	},
}
