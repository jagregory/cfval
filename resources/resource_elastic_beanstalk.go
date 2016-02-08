package resources

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
				Required: true,
			},

			"Description": Schema{
				Type: ValueString,
			},

			"SourceBundle": Schema{
				Required: true,
				Type: NestedResource{
					Description: "Elastic Beanstalk SourceBundle",
					Properties: Properties{
						"S3Bucket": Schema{
							Type:     ValueString,
							Required: true,
						},

						"S3Key": Schema{
							Type:     ValueString,
							Required: true,
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
			Required: true,
		},

		"OptionName": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Value": Schema{
			Type:     ValueString,
			Required: true,
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
				Required: true,
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
				Required: true,
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
				Type:  resourceTag,
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
