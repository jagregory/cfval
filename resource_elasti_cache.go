package main

func cacheCluster() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::CacheCluster",
		Properties: map[string]Schema{
			"AutoMinorVersionUpgrade": Schema{Type: TypeBool},
			"AZMode":                  Schema{Type: TypeString},
			"CacheNodeType":           Schema{Type: TypeString, Required: true},
			"CacheParameterGroupName": Schema{Type: TypeString},
			"CacheSecurityGroupNames": ArrayOf(Schema{Type: TypeString}),
			"CacheSubnetGroupName":    Schema{Type: TypeString},
			"ClusterName":             Schema{Type: TypeString},
			"Engine":                  Schema{Type: TypeString, Required: true},
			"EngineVersion":           Schema{Type: TypeString},
			"NotificationTopicArn":    Schema{Type: TypeString},
			"NumCacheNodes":           Schema{Type: TypeString, Required: true},
			"Port":                    Schema{Type: TypeInteger},
			"PreferredAvailabilityZone":  Schema{Type: TypeString},
			"PreferredAvailabilityZones": ArrayOf(Schema{Type: TypeString}),
			"PreferredMaintenanceWindow": Schema{Type: TypeString},
			"SnapshotArns":               ArrayOf(Schema{Type: TypeString}),
			"SnapshotName":               Schema{Type: TypeString},
			"SnapshotRetentionLimit":     Schema{Type: TypeInteger},
			"SnapshotWindow":             Schema{Type: TypeString},
			"Tags":                       ArrayOf(ResourceTag),
			"VpcSecurityGroupIds":        ArrayOf(Schema{Type: TypeString}),
		},
	}
}

func subnetGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{
			"Description": Schema{Type: TypeString, Required: true},
			"SubnetIds":   Required(ArrayOf(Schema{Type: TypeString, Required: true})),
		},
	}
}
