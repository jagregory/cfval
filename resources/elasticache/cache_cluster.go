package elasticache

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func CacheCluster() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::CacheCluster",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AutoMinorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"AZMode": Schema{
				Type: azMode,
			},

			"CacheNodeType": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"CacheParameterGroupName": Schema{
				Type: ValueString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:  ValueString,
				Array: true,
			},

			"CacheSubnetGroupName": Schema{
				Type: ValueString,
			},

			"ClusterName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"NotificationTopicArn": Schema{
				Type: ValueString,
			},

			"NumCacheNodes": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Port": Schema{
				Type: ValueNumber,
			},

			"PreferredAvailabilityZone": Schema{
				Type: ValueString,
			},

			"PreferredAvailabilityZones": Schema{
				Type:  ValueString,
				Array: true,
			},

			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"SnapshotArns": Schema{
				Type:  ValueString,
				Array: true,
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
				Type:  common.ResourceTag,
				Array: true,
			},

			"VpcSecurityGroupIds": Schema{
				Type:  ValueString,
				Array: true,
			},
		},
	}
}
