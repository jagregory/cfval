package elasti_cache

import (
	"strconv"
	"strings"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

func automaticFailoverEnabledValidation(property Schema, value interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	if version, found := self.PropertyValue("EngineVersion"); found {
		if versionNumber, err := strconv.ParseFloat(version.(string), 64); err == nil {
			if versionNumber < 2.8 {
				return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("EngineVersion must be 2.8 or higher for Automatic Failover", ctx.Path)}
			}
		}
	}

	if nodeType, found := self.PropertyValue("CacheNodeType"); found {
		split := strings.Split(nodeType.(string), ".")
		if split[1] == "t1" || split[1] == "t2" {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("CacheNodeType must not be T1 or T2 Automatic Failover", ctx.Path)}
		}
	}

	return reporting.ValidateOK, nil
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
var ReplicationGroup = Resource{
	AwsType: "AWS::ElastiCache::ReplicationGroup",

	Attributes: map[string]Schema{
		"PrimaryEndPoint.Address": Schema{
			Type: ValueString,
		},

		"PrimaryEndPoint.Port": Schema{
			Type: ValueNumber,
		},

		"ReadEndPoint.Addresses": Schema{
			Type: ValueString,
		},

		"ReadEndPoint.Ports": Schema{
			Type: ValueString,
		},

		"ReadEndPoint.Addresses.List": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ReadEndPoint.Ports.List": Schema{
			Type:  ValueNumber,
			Array: true,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: map[string]Schema{
		"AutomaticFailoverEnabled": Schema{
			Type:    ValueBool,
			Default: true,

			// You cannot enable automatic failover for Redis versions earlier than 2.8.6 or for T1 and T2 cache node types.
			ValidateFunc: automaticFailoverEnabledValidation,
		},

		// Currently, this property isn't used by ElastiCache.
		"AutoMinorVersionUpgrade": Schema{
			Type: ValueBool,
		},

		"CacheNodeType": Schema{
			Type:     cacheNodeType,
			Required: constraints.Always,
		},

		"CacheParameterGroupName": Schema{
			Type: ValueString,
		},

		"CacheSecurityGroupNames": Schema{
			Type:      cacheSecurityGroupName,
			Array:     true,
			Conflicts: constraints.PropertyExists("SecurityGroupIds"),
		},

		"CacheSubnetGroupName": Schema{
			Type: ValueString,
		},

		"Engine": Schema{
			Type:         engine,
			Required:     constraints.Always,
			ValidateFunc: SingleValueValidate("redis"),
		},

		"EngineVersion": Schema{
			Type: ValueString,
		},

		"NotificationTopicArn": Schema{
			Type: ValueString,
		},

		// If automatic failover is enabled, you must specify a value greater than 1.
		"NumCacheClusters": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
			ValidateFunc: func(prop Schema, value interface{}, self constraints.CurrentResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
				if val, ok := self.PropertyValue("AutomaticFailoverEnabled"); ok && val.(bool) == true {
					if value.(float64) <= 1 {
						return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Must be greater than 1 if automatic failover is enabled", ctx.Path)}
					}
				}

				return reporting.ValidateOK, nil
			},
		},

		"Port": Schema{
			Type: ValueNumber,
		},

		"PreferredCacheClusterAZs": Schema{
			Type:  AvailabilityZone,
			Array: true,
		},

		// Use the following format to specify a time range: ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). For example, you can specify sun:22:00-sun:23:30 for Sunday from 10 PM to 11:30 PM.
		"PreferredMaintenanceWindow": Schema{
			Type: ValueString,
		},

		"ReplicationGroupDescription": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SecurityGroupIds": Schema{
			Type:      SecurityGroupID,
			Array:     true,
			Conflicts: constraints.PropertyExists("CacheSecurityGroupNames"),
		},

		// A single-element string list that specifies an ARN of a Redis .rdb snapshot file that is stored in Amazon Simple Storage Service (Amazon S3). The snapshot file populates the node group. The Amazon S3 object name in the ARN cannot contain commas. For example, you can specify arn:aws:s3:::my_bucket/snapshot1.rdb.
		"SnapshotArns": Schema{
			Type:  ValueString,
			Array: true,
		},

		"SnapshotRetentionLimit": Schema{
			Type: ValueNumber,
		},

		// The time range (in UTC) when ElastiCache takes a daily snapshot of your node group. For example, you can specify 05:00-09:00.
		"SnapshotWindow": Schema{
			Type: ValueString,
		},
	},
}
