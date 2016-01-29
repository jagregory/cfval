package resources

import . "github.com/jagregory/cfval/schema"

func CacheCluster() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::CacheCluster",
		Properties: map[string]Schema{
			"AutoMinorVersionUpgrade": Schema{Type: TypeBool},
			"AZMode":                  Schema{Type: TypeString},
			"CacheNodeType":           Schema{Type: TypeString, Required: true},
			"CacheParameterGroupName": Schema{Type: TypeString},
			"CacheSecurityGroupNames": Schema{Type: TypeString, Array: true},
			"CacheSubnetGroupName":    Schema{Type: TypeString},
			"ClusterName":             Schema{Type: TypeString},
			"Engine":                  Schema{Type: TypeString, Required: true},
			"EngineVersion":           Schema{Type: TypeString},
			"NotificationTopicArn":    Schema{Type: TypeString},
			"NumCacheNodes":           Schema{Type: TypeString, Required: true},
			"Port":                    Schema{Type: TypeInteger},
			"PreferredAvailabilityZone":  Schema{Type: TypeString},
			"PreferredAvailabilityZones": Schema{Type: TypeString, Array: true},
			"PreferredMaintenanceWindow": Schema{Type: TypeString},
			"SnapshotArns":               Schema{Type: TypeString, Array: true},
			"SnapshotName":               Schema{Type: TypeString},
			"SnapshotRetentionLimit":     Schema{Type: TypeInteger},
			"SnapshotWindow":             Schema{Type: TypeString},
			"Tags":                       Schema{Type: resourceTag, Array: true},
			"VpcSecurityGroupIds":        Schema{Type: TypeString, Array: true},
		},
	}
}

func SubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{Type: TypeString, Required: true},
			"SubnetIds":   Schema{Type: TypeString, Required: true, Array: true},
		},
	}
}
