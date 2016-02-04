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
		Properties: map[string]Schema{
			"AutoMinorVersionUpgrade": Schema{
				Type: TypeBool,
			},

			"AZMode": Schema{
				Type:         TypeString,
				ValidateFunc: EnumValidate("single-az", "cross-az"),
			},

			"CacheNodeType": Schema{
				Type:     TypeString,
				Required: true,
			},

			"CacheParameterGroupName": Schema{
				Type: TypeString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:  TypeString,
				Array: true,
			},

			"CacheSubnetGroupName": Schema{
				Type: TypeString,
			},

			"ClusterName": Schema{
				Type: TypeString,
			},

			"Engine": Schema{
				Type:     TypeString,
				Required: true,
			},

			"EngineVersion": Schema{
				Type: TypeString,
			},

			"NotificationTopicArn": Schema{
				Type: TypeString,
			},

			"NumCacheNodes": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Port": Schema{
				Type: TypeInteger,
			},

			"PreferredAvailabilityZone": Schema{
				Type: TypeString,
			},

			"PreferredAvailabilityZones": Schema{
				Type:  TypeString,
				Array: true,
			},

			"PreferredMaintenanceWindow": Schema{
				Type: TypeString,
			},

			"SnapshotArns": Schema{
				Type:  TypeString,
				Array: true,
			},

			"SnapshotName": Schema{
				Type: TypeString,
			},

			"SnapshotRetentionLimit": Schema{
				Type: TypeInteger,
			},

			"SnapshotWindow": Schema{
				Type: TypeString,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"VpcSecurityGroupIds": Schema{
				Type:  TypeString,
				Array: true,
			},
		},
	}
}

// You cannot enable automatic failover for Redis versions earlier than 2.8.6 or for T1 and T2 cache node types.
func automaticFailoverEnabled(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if version, found := tr.Properties["EngineVersion"]; found {
		if versionNumber, err := strconv.ParseFloat(version.(string), 64); err == nil {
			if versionNumber < 2.8 {
				return false, []reporting.Failure{reporting.NewFailure("EngineVersion must be 2.8 or higher for Automatic Failover", context)}
			}
		}
	}

	if nodeType, found := tr.Properties["CacheNodeType"]; found {
		split := strings.Split(nodeType.(string), ".")
		if split[1] == "t1" || split[1] == "t2" {
			return false, []reporting.Failure{reporting.NewFailure("CacheNodeType must not be T1 or T2 Automatic Failover", context)}
		}
	}

	return true, nil
}

func ReplicationGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::ReplicationGroup",
		Properties: map[string]Schema{
			"AutomaticFailoverEnabled": Schema{
				Type:         TypeBool,
				Default:      true,
				ValidateFunc: automaticFailoverEnabled,
			},

			// Currently, this property isn't used by ElastiCache.
			"AutoMinorVersionUpgrade": Schema{
				Type: TypeBool,
			},

			"CacheNodeType": Schema{
				Type:     TypeString,
				Required: true,
				ValidateFunc: EnumValidate("cache.t2.micro", "cache.t2.small", "cache.t2.medium",
					"cache.m3.medium", "cache.m3.large", "cache.m3.xlarge", "cache.m3.2xlarge",
					"cache.t1.micro", "cache.m1.small", "cache.m1.medium", "cache.m1.large",
					"cache.m1.xlarge", "cache.c1.xlarge", "cache.r3.large", "cache.r3.xlarge",
					"cache.r3.2xlarge", "cache.r3.4xlarge", "cache.r3.8xlarge", "cache.m2.xlarge",
					"cache.m2.2xlarge", "cache.m2.4xlarge"),
			},

			"CacheParameterGroupName": Schema{
				Type: TypeString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:      TypeString,
				Array:     true,
				Conflicts: []string{"SecurityGroupIds"},
			},

			"CacheSubnetGroupName": Schema{
				Type: TypeString,
			},

			"Engine": Schema{
				Type:         TypeString,
				Required:     true,
				ValidateFunc: EnumValidate("redis"),
			},

			"EngineVersion": Schema{
				Type: TypeString,
			},

			"NotificationTopicArn": Schema{
				Type: TypeString,
			},

			// If automatic failover is enabled, you must specify a value greater than 1.
			"NumCacheClusters": Schema{
				Type:     TypeInteger,
				Required: true,
			},

			"Port": Schema{
				Type: TypeInteger,
			},

			"PreferredCacheClusterAZs": Schema{
				Type:         TypeString,
				Array:        true,
				ValidateFunc: availabilityZone,
			},

			// Use the following format to specify a time range: ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). For example, you can specify sun:22:00-sun:23:30 for Sunday from 10 PM to 11:30 PM.
			"PreferredMaintenanceWindow": Schema{
				Type: TypeString,
			},

			"ReplicationGroupDescription": Schema{
				Type:     TypeString,
				Required: true,
			},

			"SecurityGroupIds": Schema{
				Type:      TypeString,
				Array:     true,
				Conflicts: []string{"CacheSecurityGroupNames"},
			},

			// A single-element string list that specifies an ARN of a Redis .rdb snapshot file that is stored in Amazon Simple Storage Service (Amazon S3). The snapshot file populates the node group. The Amazon S3 object name in the ARN cannot contain commas. For example, you can specify arn:aws:s3:::my_bucket/snapshot1.rdb.
			"SnapshotArns": Schema{
				Type:  TypeString,
				Array: true,
			},

			"SnapshotRetentionLimit": Schema{
				Type: TypeInteger,
			},

			// The time range (in UTC) when ElastiCache takes a daily snapshot of your node group. For example, you can specify 05:00-09:00.
			"SnapshotWindow": Schema{
				Type: TypeString,
			},
		},
	}
}

func SubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{
				Type:     TypeString,
				Required: true,
			},

			"SubnetIds": Schema{
				Type:     TypeString,
				Required: true,
				Array:    true,
			},
		},
	}
}
