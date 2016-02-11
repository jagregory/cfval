package resources

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

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

var optionsSettings = NestedResource{
	Description: "Elastic Beanstalk OptionSettings",
	Properties: Properties{
		"Namespace": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OptionName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

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
