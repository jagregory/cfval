package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html
var NetworkInterfaceAttachment = Resource{
	AwsType: "AWS::EC2::NetworkInterfaceAttachment",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DeleteOnTermination": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"DeviceIndex": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"InstanceId": Schema{
			Type:     InstanceID,
			Required: constraints.Always,
		},

		"NetworkInterfaceId": Schema{
			Type:     NetworkInterfaceID,
			Required: constraints.Always,
		},
	},
}
