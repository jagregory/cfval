package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
func EIPAssociation() Resource {
	return Resource{
		AwsType: "AWS::EC2::EIPAssociation",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AllocationId": Schema{
				Type: ValueString,
			},

			"EIP": Schema{
				Type: IPAddress,
			},

			"InstanceId": Schema{
				Type: InstanceID,
			},

			"NetworkInterfaceId": Schema{
				Type: ValueString,
			},

			"PrivateIpAddress": Schema{
				Type: IPAddress,
			},
		},
	}
}
