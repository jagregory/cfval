package s3

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
var websiteConfigurationRedirectAllRequestsTo = NestedResource{
	Description: "S3 Website Configuration Redirect All Requests To",
	Properties: Properties{
		"HostName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Protocol": Schema{
			Type: EnumValue{
				Description: "Redirect Protocol",
				Options:     []string{"http", "https"},
			},
		},
	},
}
