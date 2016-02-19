package cloud_formation

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
var WaitCondition = Resource{
	AwsType: "AWS::CloudFormation::WaitCondition",

	Attributes: map[string]Schema{
		"Data": Schema{
			Type: JSON,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{},
}
