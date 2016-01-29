package resources

import . "github.com/jagregory/cfval/schema"

func Application() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::Application",
		Properties: map[string]Schema{
			"ApplicationName": Schema{Type: TypeString},
			"Description":     Schema{Type: TypeString},
		},
	}
}

func ApplicationVersion() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ApplicationVersion",
		Properties: map[string]Schema{
			"ApplicationName": Schema{Type: TypeString, Required: true},
			"Description":     Schema{Type: TypeString},
			"SourceBundle": Schema{
				Required: true,
				Type: Resource{
					AwsType: "Elastic Beanstalk SourceBundle",
					Properties: map[string]Schema{
						"S3Bucket": Schema{Type: TypeString, Required: true},
						"S3Key":    Schema{Type: TypeString, Required: true},
					},
				},
			},
		},
	}
}

var optionsSettings = Resource{
	AwsType: "Elastic Beanstalk OptionSettings",
	Properties: map[string]Schema{
		"Namespace":  Schema{Type: TypeString, Required: true},
		"OptionName": Schema{Type: TypeString, Required: true},
		"Value":      Schema{Type: TypeString, Required: true},
	},
}

func ConfigurationTemplate() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::ConfigurationTemplate",
		Properties: map[string]Schema{
			"ApplicationName": Schema{Type: TypeString, Required: true},
			"Description":     Schema{Type: TypeString},
			// "EnvironmentId": Schema{Type:TypeString},
			"OptionSettings":    ArrayOf(Schema{Type: optionsSettings}),
			"SolutionStackName": Schema{Type: TypeString},
			// "SourceConfiguration": Schema{Type:...},
		},
	}
}

func Environment() Resource {
	return Resource{
		AwsType: "AWS::ElasticBeanstalk::Environment",
		Properties: map[string]Schema{
			"ApplicationName":   Schema{Type: TypeString, Required: true},
			"CNAMEPrefix":       Schema{Type: TypeString},
			"Description":       Schema{Type: TypeString},
			"EnvironmentName":   Schema{Type: TypeString},
			"OptionSettings":    ArrayOf(Schema{Type: optionsSettings}),
			"SolutionStackName": Schema{Type: TypeString},
			"Tags":              ArrayOf(resourceTag),
			"TemplateName":      Schema{Type: TypeString},
			// "Tier": Schema{...}
			"VersionLabel": Schema{Type: TypeString},
		},
	}
}
