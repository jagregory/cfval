package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
func LaunchConfiguration() Resource {
	return Resource{
		AwsType: "AWS::AutoScaling::LaunchConfiguration",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AssociatePublicIpAddress": Schema{
				Type: ValueBool,
			},

			"BlockDeviceMappings": Schema{
				Type:  autoScalingBlockDeviceMapping,
				Array: true,
			},

			"ClassicLinkVPCId": Schema{
				Type: ValueString,
			},

			"ClassicLinkVPCSecurityGroups": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.PropertyExists("ClassicLinkVPCId"),
			},

			"EbsOptimized": Schema{
				Type:    ValueBool,
				Default: false,
			},

			"IamInstanceProfile": Schema{
				Type:         ValueString,
				ValidateFunc: StringLengthValidate(1, 1600),
			},

			"ImageId": Schema{
				Type: ValueString,
			},

			"InstanceId": Schema{
				Type: ValueString,
			},

			"InstanceMonitoring": Schema{
				Type:    ValueBool,
				Default: true,
			},

			"InstanceType": Schema{
				Type: ValueString,
			},

			"KernelId": Schema{
				Type: ValueString,
			},

			"KeyName": Schema{
				Type: KeyName,
			},

			// TODO: If you specify this property, you must specify at least one subnet in the VPCZoneIdentifier property of the AWS::AutoScaling::AutoScalingGroup resource.
			// This will require some reverse lookups from this resource to any which use it: not supported yet.
			"PlacementTenancy": Schema{
				Type: placementTenancy,
			},

			"RamDiskId": Schema{
				Type: ValueString,
			},

			"SecurityGroups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"SpotPrice": Schema{
				Type: ValueString,
			},

			"UserData": Schema{
				Type: ValueString,
			},
		},
	}
}
