package cloud_formation

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
var WaitConditionHandle = Resource{
	AwsType: "AWS::CloudFormation::WaitConditionHandle",

	// URL
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{},
}
