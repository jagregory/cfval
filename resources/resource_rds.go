package resources

import . "github.com/jagregory/cfval/schema"

func DBInstance() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBInstance",

		// ID
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"AllocatedStorage": Schema{
				Type: TypeString,
			},

			"AllowMajorVersionUpgrade": Schema{
				Type: TypeBool,
			},

			"AutoMinorVersionUpgrade": Schema{
				Type: TypeBool,
			},

			"AvailabilityZone": Schema{
				Type: TypeString,
			},

			"BackupRetentionPeriod": Schema{
				Type: TypeString,
			},

			"CharacterSetName": Schema{
				Type: TypeString,
			},

			"DBClusterIdentifier": Schema{
				Type: TypeString,
			},

			"DBInstanceClass": Schema{
				Type:     TypeString,
				Required: true,
			},

			"DBInstanceIdentifier": Schema{
				Type: TypeString,
			},

			"DBName": Schema{
				Type: TypeString,
			},

			"DBParameterGroupName": Schema{
				Type: TypeString,
			},

			"DBSecurityGroups": Schema{
				Type:  TypeString,
				Array: true,
			},

			"DBSnapshotIdentifier": Schema{
				Type: TypeString,
			},

			"DBSubnetGroupName": Schema{
				Type: TypeString,
			},

			"Engine": Schema{
				Type: TypeString,
			},

			"EngineVersion": Schema{
				Type: TypeString,
			},

			"Iops": Schema{
				Type: TypeInteger,
			},

			"KmsKeyId": Schema{
				Type: TypeString,
			},

			"LicenseModel": Schema{
				Type: TypeString,
			},

			"MasterUsername": Schema{
				Type: TypeString,
			},

			"MasterUserPassword": Schema{
				Type: TypeString,
			},

			"MultiAZ": Schema{
				Type: TypeBool,
			},

			"OptionGroupName": Schema{
				Type: TypeString,
			},

			"Port": Schema{
				Type: TypeString,
			},

			"PreferredBackupWindow": Schema{
				Type: TypeString,
			},

			"PreferredMaintenanceWindow": Schema{
				Type: TypeString,
			},

			"PubliclyAccessible": Schema{
				Type: TypeBool,
			},

			"SourceDBInstanceIdentifier": Schema{
				Type: TypeString,
			},

			"StorageEncrypted": Schema{
				Type: TypeBool,
			},

			"StorageType": Schema{
				Type: TypeString,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			"VPCSecurityGroups": Schema{
				Type:  TypeString,
				Array: true,
			},
		},
	}
}

func DBSubnetGroup() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBSubnetGroup",
		Properties: map[string]Schema{
			"DBSubnetGroupDescription": Schema{Type: TypeString, Required: true},
			"SubnetIds":                Schema{Type: TypeString, Required: true, Array: true},
			"Tags":                     Schema{Type: resourceTag, Array: true},
		},
	}
}
