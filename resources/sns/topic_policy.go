package sns

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
func TopicPolicy() Resource {
	return Resource{
		AwsType: "AWS::SNS::TopicPolicy",

		Properties: Properties{
			"PolicyDocument": Schema{
				Type:     JSON,
				Required: constraints.Always,
			},

			"Topics": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
