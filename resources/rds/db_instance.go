package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html
func DBInstance() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBInstance",

		// ID
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AllocatedStorage": Schema{
				Type: ValueString,
				// TODO: If any value is used in the Iops parameter, AllocatedStorage
				//			 must be at least 100 GB, which corresponds to the minimum Iops
				//       value of 1000. If Iops is increased (in 1000 IOPS increments),
				//       then AllocatedStorage must also be increased (in 100 GB
				//       increments) correspondingly.
				Required:  constraints.PropertyNotExists("DBClusterIdentifier"),
				Conflicts: constraints.PropertyExists("DBClusterIdentifier"),
			},

			"AllowMajorVersionUpgrade": Schema{
				Type: ValueBool,
				// TODO: This parameter must be set to true when you specify an
				//       EngineVersion that differs from the DB instance's current
				//       major version.
			},

			"AutoMinorVersionUpgrade": Schema{
				Type:    ValueBool,
				Default: true,
			},

			"AvailabilityZone": Schema{
				Type:      AvailabilityZone,
				Conflicts: constraints.PropertyIs("MultiAZ", true),
			},

			"BackupRetentionPeriod": Schema{
				Type: ValueString,
				Conflicts: constraints.Any{
					constraints.PropertyExists("DBSnapshotIdentifier"),
					constraints.PropertyExists("SourceDBInstanceIdentifier"),
				},
			},

			"CharacterSetName": Schema{
				Type: ValueString,
			},

			"DBClusterIdentifier": Schema{
				Type: ValueString,
				// TODO: The identifier of an existing DB cluster that this instance
				// 			 will be associated with. If you specify this property, specify
				//       aurora for the Engine property and do not specify any of the
				//       following properties: AllocatedStorage, CharacterSetName,
				//       DBSecurityGroups, SourceDBInstanceIdentifier, and StorageType.
			},

			"DBInstanceClass": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"DBInstanceIdentifier": Schema{
				Type: ValueString,
			},

			"DBName": Schema{
				Type: ValueString,
			},

			"DBParameterGroupName": Schema{
				Type: dbParameterGroupName,
			},

			"DBSecurityGroups": Schema{
				Type:      dbSecurityGroupName,
				Array:     true,
				Conflicts: constraints.PropertyExists("VPCSecurityGroups"),
			},

			"DBSnapshotIdentifier": Schema{
				Type: ValueString,
				Conflicts: constraints.Any{
					constraints.PropertyExists("MasterUsername"),
					constraints.PropertyExists("MasterUserPassword"),
				},
			},

			"DBSubnetGroupName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type: EnumValue{
					Description: "DB Instance Engine",
					Options: []string{
						"MySQL",
						"mariadb",
						"oracle-se1",
						"oracle-se",
						"oracle-ee",
						"sqlserver-ee",
						"sqlserver-se",
						"sqlserver-ex",
						"sqlserver-web",
						"postgres",
						"aurora",
					},
				},
				Required: constraints.PropertyNotExists("DBSnapshotIdentifier"),
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"Iops": Schema{
				Type:     ValueNumber,
				Required: constraints.PropertyIs("StorageType", "io1"),
			},

			"KmsKeyId": Schema{
				Type: ValueString,
				Conflicts: constraints.Any{
					constraints.PropertyExists("DBSnapshotIdentifier"),
					constraints.PropertyExists("SourceDBInstanceIdentifier"),
				},
				// TODO: If you specify this property, you must set the StorageEncrypted
				//			 property to true.
			},

			"LicenseModel": Schema{
				Type: ValueString,
			},

			"MasterUsername": Schema{
				Type: ValueString,
				Conflicts: constraints.Any{
					constraints.PropertyExists("SourceDBInstanceIdentifier"),
					constraints.PropertyExists("DBSnapshotIdentifier"),
				},
			},

			"MasterUserPassword": Schema{
				Type:      ValueString,
				Conflicts: constraints.PropertyExists("SourceDBInstanceIdentifier"),
			},

			"MultiAZ": Schema{
				Type: ValueBool,
			},

			"OptionGroupName": Schema{
				Type: ValueString,
			},

			"Port": Schema{
				Type: ValueString,
			},

			"PreferredBackupWindow": Schema{
				Type: ValueString,
			},

			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"PubliclyAccessible": Schema{
				Type: ValueBool,
			},

			"SourceDBInstanceIdentifier": Schema{
				Type: ValueString,
				Conflicts: constraints.Any{
					constraints.PropertyIs("MultiAZ", true),
					constraints.PropertyExists("DBSnapshotIdentifier"),
					constraints.PropertyExists("BackupRetentionPeriod"),
					constraints.PropertyExists("DBName"),
					constraints.PropertyExists("MasterUsername"),
					constraints.PropertyExists("MasterUserPassword"),
					constraints.PropertyExists("PreferredBackupWindow"),
					constraints.PropertyIs("Engine", "aurora"),
				},
			},

			"StorageEncrypted": Schema{
				Type: ValueBool,
				Conflicts: constraints.Any{
					constraints.PropertyExists("DBSnapshotIdentifier"),
					constraints.PropertyExists("SourceDBInstanceIdentifier"),
				},
				Required: constraints.PropertyExists("KmsKeyId"),
			},

			"StorageType": Schema{
				Type: EnumValue{
					Description: "DB Instance Storage Type",

					Options: []string{"standard", "io1", "gp2"},
				},
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VPCSecurityGroups": Schema{
				Type:      SecurityGroupID,
				Array:     true,
				Conflicts: constraints.PropertyExists("DBSubnetGroups"),
			},
		},
	}
}
