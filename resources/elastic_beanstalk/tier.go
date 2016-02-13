package elastic_beanstalk

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html
var tier = NestedResource{
	Description: "Elastic Beanstalk Environment Tier",
	Properties: Properties{
		"Name": Schema{
			Type: EnumValue{
				Description: "Tier Name",
				Options:     []string{"WebServer", "Worker"},
			},
		},

		"Type": Schema{
			Type: EnumValue{
				Description: "Tier Type",
				Options:     []string{"Standard", "SQS", "HTTP"},
			},
			ValidateFunc: func(prop Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Reports) {
				name, _ := self.Property("Name")

				if name == "WebServer" && value != "Standard" {
					return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Must be Standard for WebServer tier", context)}
				}

				return reporting.ValidateOK, nil
			},
		},

		"Version": Schema{
			Type: ValueString,
		},
	},
}
