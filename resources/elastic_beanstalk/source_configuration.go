package elastic_beanstalk

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html
var sourceConfiguration = NestedResource{
	Description: "Elastic Beanstalk SourceConfiguration",
	Properties: Properties{
		"ApplicationName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TemplateName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
