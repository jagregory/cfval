package auto_scaling

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
var AutoScalingGroup = Resource{
	AwsType: "AWS::AutoScaling::AutoScalingGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AvailabilityZones": Schema{
			Array:    true,
			Type:     AvailabilityZone,
			Required: constraints.PropertyNotExists("VPCZoneIdentifier"),
		},

		"Cooldown": Schema{
			Type: ValueString,
		},

		"DesiredCapacity": Schema{
			Type: ValueString,
		},

		"HealthCheckGracePeriod": Schema{
			Type: ValueNumber,
		},

		"HealthCheckType": Schema{
			Type: healthCheckType,
		},

		"InstanceId": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("LaunchConfigurationName"),
		},

		"LaunchConfigurationName": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("InstanceId"),
		},

		"LoadBalancerNames": Schema{
			Type:  ValueString,
			Array: true,
		},

		"MaxSize": Schema{
			Type: ValueString,
		},

		"MetricsCollection": Schema{
			Type:  metricsCollection,
			Array: true,
		},

		"MinSize": Schema{
			Type: ValueString,
		},

		"NotificationConfigurations": Schema{
			Type:  notificationConfiguration,
			Array: true,
		},

		"PlacementGroup": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type:  autoScalingTag,
			Array: true,
		},

		"TerminationPolicies": Schema{
			Type:  ValueString,
			Array: true,
		},

		"VPCZoneIdentifier": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.PropertyNotExists("AvailabilityZones"),
		},
	},
}
