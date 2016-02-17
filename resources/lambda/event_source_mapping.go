package lambda

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
func EventSourceMapping() Resource {
	return Resource{
		AwsType: "AWS::Lambda::EventSourceMapping",

		Properties: Properties{
			"BatchSize": Schema{
				Type:         ValueNumber,
				ValidateFunc: IntegerRangeValidate(1, 10000),
			},

			"Enabled": Schema{
				Type: ValueBool,
			},

			"EventSourceArn": Schema{
				Type:     ValueString,
				Required: constraints.Always,
				ValidateFunc: RegexpValidate(
					`^arn:aws:([a-zA-Z0-9\-])+:([a-z]{2}-[a-z]+-\d{1})?:(\d{12})?:(.*)$`,
					"The Amazon Resource Name (ARN) of the Amazon Kinesis stream that is the source of events.",
				),
			},

			"FunctionName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
				ValidateFunc: RegexpValidate(
					`^(arn:aws:lambda:)?([a-z]{2}-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`,
					"You can specify the function name (for example, Thumbnail) or you can specify Amazon Resource Name (ARN) of the function (for example, arn:aws:lambda:us-west-2:account-id:function:ThumbNail).",
				),
			},

			"StartingPosition": Schema{
				Type: EnumValue{
					Description: "StartingPosition",
					Options:     []string{"TRIM_HORIZON", "LATEST"},
				},
				Required: constraints.Always,
			},
		},
	}
}
