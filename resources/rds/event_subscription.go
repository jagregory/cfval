package rds

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
var EventSubscription = Resource{
	AwsType: "AWS::RDS::EventSubscription",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Enabled": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"EventCategories": Schema{
			Type: Multiple(ValueString),
		},

		"SnsTopicArn": Schema{
			Type:     ARN,
			Required: constraints.Always,
		},

		"SourceIds": Schema{
			Type: Multiple(ValueString),
		},

		"SourceType": Schema{
			Type: ValueString,
			Required: constraints.Any{
				constraints.PropertyExists("SourceIds"),
				constraints.PropertyExists("EventCategories"),
			},
		},
	},
}
