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
			Type:     Multiple(AvailabilityZone),
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
			Type:     InstanceID,
			Required: constraints.PropertyNotExists("LaunchConfigurationName"),
		},

		"LaunchConfigurationName": Schema{
			Type:     ValueString,
			Required: constraints.PropertyNotExists("InstanceId"),
		},

		"LoadBalancerNames": Schema{
			Type: Multiple(ValueString),
		},

		"MaxSize": Schema{
			Type: ValueString,
		},

		"MetricsCollection": Schema{
			Type: Multiple(metricsCollection),
		},

		"MinSize": Schema{
			Type: ValueString,
		},

		"NotificationConfigurations": Schema{
			Type: Multiple(notificationConfiguration),
		},

		"PlacementGroup": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type: Multiple(autoScalingTag),
		},

		"TerminationPolicies": Schema{
			Type: Multiple(ValueString),
		},

		"VPCZoneIdentifier": Schema{
			Type:     Multiple(SubnetID),
			Required: constraints.PropertyNotExists("AvailabilityZones"),
		},
	},
}
