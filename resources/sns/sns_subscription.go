package sns

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
var snsSubscription = NestedResource{
	Description: "SNS Subscription Property",
	Properties: Properties{
		// TODO: richer validation: http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html
		"Endpoint": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Protocol": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
