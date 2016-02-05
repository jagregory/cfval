package resources

import . "github.com/jagregory/cfval/schema"

func Application() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::Application",

		// Name
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"ApplicationName": Schema{
				Type: TypeString,
			},

			"Description": Schema{
				Type: TypeString,
			},
		},
	}
}

func ApplicationVersion() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ApplicationVersion",

		// Name
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"ApplicationName": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Description": Schema{
				Type: TypeString,
			},

			"SourceBundle": Schema{
				Required: true,
				Type: Resource{
					AwsType: "Elastic Beanstalk SourceBundle",
					Properties: map[string]Schema{
						"S3Bucket": Schema{
							Type:     TypeString,
							Required: true,
						},

						"S3Key": Schema{
							Type:     TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

var optionsSettings = Resource{
	AwsType: "Elastic Beanstalk OptionSettings",
	Properties: map[string]Schema{
		"Namespace": Schema{
			Type:     TypeString,
			Required: true,
		},

		"OptionName": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Value": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

func ConfigurationTemplate() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ConfigurationTemplate",

		// Name
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"ApplicationName": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Description": Schema{
				Type: TypeString,
			},

			// "EnvironmentId": Schema{Type:TypeString},

			"OptionSettings": Schema{
				Type:  optionsSettings,
				Array: true,
			},

			"SolutionStackName": Schema{
				Type: TypeString,
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
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"ApplicationName": Schema{
				Type:     TypeString,
				Required: true,
			},

			"CNAMEPrefix": Schema{
				Type: TypeString,
			},

			"Description": Schema{
				Type: TypeString,
			},

			"EnvironmentName": Schema{
				Type: TypeString,
			},

			"OptionSettings": Schema{
				Type:  optionsSettings,
				Array: true,
			},

			"SolutionStackName": Schema{
				Type: TypeString,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"TemplateName": Schema{
				Type: TypeString,
			},

			// "Tier": Schema{...}

			"VersionLabel": Schema{
				Type: TypeString,
			},
		},
	}
}
