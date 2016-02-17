package logs

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
func LogStream() Resource {
	return Resource{
		AwsType: "AWS::Logs::LogStream",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"LogGroupName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"LogStreamName": Schema{
				Type: ValueString,
			},
		},
	}
}
