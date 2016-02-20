package ec2

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
var EIPAssociation = Resource{
	AwsType: "AWS::EC2::EIPAssociation",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AllocationId": Schema{
			Type: AllocationID,
		},

		"EIP": Schema{
			Type: IPAddress,
		},

		"InstanceId": Schema{
			Type: InstanceID,
		},

		"NetworkInterfaceId": Schema{
			Type: NetworkInterfaceID,
		},

		"PrivateIpAddress": Schema{
			Type: IPAddress,
		},
	},
}
