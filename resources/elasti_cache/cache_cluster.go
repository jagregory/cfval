package elasti_cache

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func azModeValidate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if str, ok := value.(string); ok {
		if availabilityZones, ok := ctx.CurrentResource().PropertyValueOrDefault("PreferredAvailabilityZones"); ok {
			if str == "cross-az" && len(availabilityZones.([]interface{})) < 2 {
				return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Cross-AZ clusters must have multiple preferred availability zones")}
			}
		}
	}

	return reporting.ValidateOK, nil
}

func numCacheNodesValidate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if engine, ok := ctx.CurrentResource().PropertyValueOrDefault("Engine"); !ok || engine.(string) == "memcached" {
		return IntegerRangeValidate(1, 20)(value, ctx)
	}

	return SingleValueValidate(float64(1))(value, ctx)
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html
var CacheCluster = Resource{
	AwsType: "AWS::ElastiCache::CacheCluster",

	Attributes: map[string]Schema{
		"ConfigurationEndpoint.Address": Schema{
			Type: ValueString,
		},

		"ConfigurationEndpoint.Port": Schema{
			Type: ValueString,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AutoMinorVersionUpgrade": Schema{
			Type: ValueBool,
		},

		"AZMode": Schema{
			Type:         azMode,
			ValidateFunc: azModeValidate,
			Default:      "single-az",
		},

		"CacheNodeType": Schema{
			Type:     cacheNodeType,
			Required: constraints.Always,
		},

		"CacheParameterGroupName": Schema{
			Type: ValueString,
		},

		"CacheSecurityGroupNames": Schema{
			Type: Multiple(cacheSecurityGroupName),
			Conflicts: constraints.Any{
				constraints.PropertyExists("CacheSubnetGroupName"),
				constraints.PropertyExists("VpcSecurityGroupIds"),
			},
		},

		"CacheSubnetGroupName": Schema{
			Type: cacheSecurityGroupName,
			Conflicts: constraints.Any{
				constraints.PropertyExists("CacheSecurityGroupNames"),
				constraints.PropertyExists("VpcSecurityGroupIds"),
			},
		},

		"ClusterName": Schema{
			Type: ValueString,
		},

		"Engine": Schema{
			Type:     engine,
			Required: constraints.Always,
		},

		"EngineVersion": Schema{
			Type: ValueString,
		},

		"NotificationTopicArn": Schema{
			Type: ARN,
		},

		"NumCacheNodes": Schema{
			Type:         ValueNumber,
			Required:     constraints.Always,
			ValidateFunc: numCacheNodesValidate,
		},

		"Port": Schema{
			Type: ValueNumber,
		},

		"PreferredAvailabilityZone": Schema{
			Type: AvailabilityZone,
		},

		"PreferredAvailabilityZones": Schema{
			Type:     Multiple(AvailabilityZone),
			Required: constraints.PropertyIs("AZMode", "cross-az"),
		},

		"PreferredMaintenanceWindow": Schema{
			Type: ValueString,
		},

		"SnapshotArns": Schema{
			Type: Multiple(ARN),
		},

		"SnapshotName": Schema{
			Type: ValueString,
		},

		"SnapshotRetentionLimit": Schema{
			Type: ValueNumber,
		},

		"SnapshotWindow": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},

		"VpcSecurityGroupIds": Schema{
			Type:      Multiple(SecurityGroupID),
			Conflicts: constraints.PropertyExists("CacheSecurityGroupNames"),
		},
	},
}
