package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
var Eip = Resource{
	AwsType: "AWS::EC2::EIP",

	Attributes: map[string]Schema{
		"AllocationId": Schema{
			Type: ValueString,
		},
	},

	// PublicIp
	ReturnValue: Schema{
		Type: IPAddress,
	},

	Properties: Properties{
		"InstanceId": Schema{
			Type: InstanceID,
		},

		"Domain": Schema{
			Type: EnumValue{
				Description: "EIP Domain",
				Options:     []string{"vpc"},
			},
		},
	},
}
