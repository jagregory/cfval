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
			ValidateFunc: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
				name, _ := ctx.CurrentResource().PropertyValueOrDefault("Name")

				if name == "WebServer" && value != "Standard" {
					return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Must be Standard for WebServer tier")}
				}

				return reporting.ValidateOK, nil
			},
		},

		"Version": Schema{
			Type: ValueString,
		},
	},
}
