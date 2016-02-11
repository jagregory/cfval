package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
func Eip() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIP",

		// PublicIp
		ReturnValue: Schema{
			Type: ValueString,
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
}
