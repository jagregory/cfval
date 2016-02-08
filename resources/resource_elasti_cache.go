package resources

import (
	"strconv"
	"strings"

	"github.com/jagregory/cfval/reporting"
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
				Type: EnumValue{[]string{"single-az", "cross-az"}},
			},

			"CacheNodeType": Schema{
				Type:     ValueString,
				Required: true,
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
				Required: true,
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"NotificationTopicArn": Schema{
				Type: ValueString,
			},

			"NumCacheNodes": Schema{
				Type:     ValueString,
				Required: true,
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
				Type:  resourceTag,
				Array: true,
			},

			"VpcSecurityGroupIds": Schema{
				Type:  ValueString,
				Array: true,
			},
		},
	}
}

// You cannot enable automatic failover for Redis versions earlier than 2.8.6 or for T1 and T2 cache node types.
var automaticFailoverEnabled FuncType = func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if version, found := self.Property("EngineVersion"); found {
		if versionNumber, err := strconv.ParseFloat(version.(string), 64); err == nil {
			if versionNumber < 2.8 {
				return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure("EngineVersion must be 2.8 or higher for Automatic Failover", context)}
			}
		}
	}

	if nodeType, found := self.Property("CacheNodeType"); found {
		split := strings.Split(nodeType.(string), ".")
		if split[1] == "t1" || split[1] == "t2" {
			return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure("CacheNodeType must not be T1 or T2 Automatic Failover", context)}
		}
	}

	return reporting.ValidateOK, nil
}

func ReplicationGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::ReplicationGroup",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AutomaticFailoverEnabled": Schema{
				Type:    automaticFailoverEnabled,
				Default: true,
			},

			// Currently, this property isn't used by ElastiCache.
			"AutoMinorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"CacheNodeType": Schema{
				Type: EnumValue{[]string{"cache.t2.micro", "cache.t2.small", "cache.t2.medium",
					"cache.m3.medium", "cache.m3.large", "cache.m3.xlarge", "cache.m3.2xlarge",
					"cache.t1.micro", "cache.m1.small", "cache.m1.medium", "cache.m1.large",
					"cache.m1.xlarge", "cache.c1.xlarge", "cache.r3.large", "cache.r3.xlarge",
					"cache.r3.2xlarge", "cache.r3.4xlarge", "cache.r3.8xlarge", "cache.m2.xlarge",
					"cache.m2.2xlarge", "cache.m2.4xlarge"}},
				Required: true,
			},

			"CacheParameterGroupName": Schema{
				Type: ValueString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:      ValueString,
				Array:     true,
				Conflicts: []string{"SecurityGroupIds"},
			},

			"CacheSubnetGroupName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type:     EnumValue{[]string{"redis"}},
				Required: true,
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
				Required: true,
			},

			"Port": Schema{
				Type: ValueNumber,
			},

			"PreferredCacheClusterAZs": Schema{
				Type:  availabilityZone,
				Array: true,
			},

			// Use the following format to specify a time range: ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). For example, you can specify sun:22:00-sun:23:30 for Sunday from 10 PM to 11:30 PM.
			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"ReplicationGroupDescription": Schema{
				Type:     ValueString,
				Required: true,
			},

			"SecurityGroupIds": Schema{
				Type:      ValueString,
				Array:     true,
				Conflicts: []string{"CacheSecurityGroupNames"},
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
}

func SubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{
				Type:     ValueString,
				Required: true,
			},

			"SubnetIds": Schema{
				Type:     ValueString,
				Required: true,
				Array:    true,
			},
		},
	}
}
