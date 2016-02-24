package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
var LoadBalancer = Resource{
	AwsType: "AWS::ElasticLoadBalancing::LoadBalancer",

	Attributes: map[string]Schema{
		"CanonicalHostedZoneName": Schema{
			Type: ValueString,
		},

		"CanonicalHostedZoneNameID": Schema{
			Type: HostedZoneID,
		},

		"DNSName": Schema{
			Type: ValueString,
		},

		"SourceSecurityGroup.GroupName": Schema{
			Type: SecurityGroupName,
		},

		"SourceSecurityGroup.OwnerAlias": Schema{
			Type: ValueString,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AccessLoggingPolicy": Schema{
			Type: accessLoggingPolicy,
		},

		"AppCookieStickinessPolicy": Schema{
			Type: Multiple(appCookieStickinessPolicy),
		},

		"AvailabilityZones": Schema{
			Type:      Multiple(AvailabilityZone),
			Conflicts: constraints.PropertyExists("Subnets"),
		},

		"ConnectionDrainingPolicy": Schema{
			Type: connectionDrainingPolicy,
		},

		"ConnectionSettings": Schema{
			Type: connectionSettings,
		},

		"CrossZone": Schema{
			Type:    ValueBool,
			Default: false,
		},

		"HealthCheck": Schema{
			Type: healthCheck,
		},

		"Instances": Schema{
			Type: Multiple(InstanceID),
		},

		"LBCookieStickinessPolicy": Schema{
			Type: lbCookieStickinessPolicy,
		},

		"LoadBalancerName": Schema{
			Type: ValueString,
		},

		"Listeners": Schema{
			Type:     Multiple(listener),
			Required: constraints.Always,
		},

		"Policies": Schema{
			Type: Multiple(policy),
		},

		"Scheme": Schema{
			Type: EnumValue{
				Description: "Load Balancer Scheme",
				Options:     []string{"internal", "internet-facing"},
				// TODO: If you specify internal, you must specify subnets to associate with the load balancer, not Availability Zones.
			},
		},

		"SecurityGroups": Schema{
			Type: Multiple(SecurityGroupID),
		},

		"Subnets": Schema{
			Type:      Multiple(SubnetID),
			Conflicts: constraints.PropertyExists("AvailabilityZones"),
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
