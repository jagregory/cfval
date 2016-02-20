package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
var Route = Resource{
	AwsType: "AWS::EC2::Route",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"DestinationCidrBlock": Schema{
			Type:     CIDR,
			Required: constraints.Always,
		},

		"GatewayId": Schema{
			Type: InternetGatewayID,
			Required: constraints.All{
				constraints.PropertyNotExists("InstanceId"),
				constraints.PropertyNotExists("NetworkInterfaceId"),
				constraints.PropertyNotExists("VpcPeeringConnectionId"),
			},
			Conflicts: constraints.Any{
				constraints.PropertyExists("InstanceId"),
				constraints.PropertyExists("NetworkInterfaceId"),
				constraints.PropertyExists("VpcPeeringConnectionId"),
			},
		},

		"InstanceId": Schema{
			Type: InstanceID,
			Required: constraints.All{
				constraints.PropertyNotExists("GatewayId"),
				constraints.PropertyNotExists("NetworkInterfaceId"),
				constraints.PropertyNotExists("VpcPeeringConnectionId"),
			},
			Conflicts: constraints.Any{
				constraints.PropertyExists("GatewayId"),
				constraints.PropertyExists("NetworkInterfaceId"),
				constraints.PropertyExists("VpcPeeringConnectionId"),
			},
		},

		"NetworkInterfaceId": Schema{
			Type: NetworkInterfaceID,
			Required: constraints.All{
				constraints.PropertyNotExists("GatewayId"),
				constraints.PropertyNotExists("InstanceId"),
				constraints.PropertyNotExists("VpcPeeringConnectionId"),
			},
			Conflicts: constraints.Any{
				constraints.PropertyExists("GatewayId"),
				constraints.PropertyExists("InstanceId"),
				constraints.PropertyExists("VpcPeeringConnectionId"),
			},
		},

		"RouteTableId": Schema{
			Type:     RouteTableID,
			Required: constraints.Always,
		},

		"VpcPeeringConnectionId": Schema{
			Type: VpcPeeringConnectionID,
			Required: constraints.All{
				constraints.PropertyNotExists("GatewayId"),
				constraints.PropertyNotExists("InstanceId"),
				constraints.PropertyNotExists("NetworkInterfaceId"),
			},
			Conflicts: constraints.Any{
				constraints.PropertyExists("GatewayId"),
				constraints.PropertyExists("InstanceId"),
				constraints.PropertyExists("NetworkInterfaceId"),
			},
		},
	},
}
