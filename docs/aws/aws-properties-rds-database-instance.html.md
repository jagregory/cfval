AWS::RDS::DBInstance
====================

The `AWS::RDS::DBInstance` type creates an Amazon RDS database instance. For detailed information about configuring RDS DB instances, see [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).

Important

If a DB instance is deleted or replaced during an update, all automated snapshots are deleted. However, manual DB snapshot are retained. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see [Prevent Updates to Stack Resources](protect-stack-resources.html "Prevent Updates to Stack Resources").

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::RDS::DBInstance",
  "Properties" :
  {
    "AllocatedStorage" : String,
    "AllowMajorVersionUpgrade" : Boolean,
    "AutoMinorVersionUpgrade" : Boolean,
    "AvailabilityZone" : String,
    "BackupRetentionPeriod" : String,
    "CharacterSetName" : String,
    "DBClusterIdentifier" : String,
    "DBInstanceClass" : String,
    "DBInstanceIdentifier" : String,
    "DBName" : String,
    "DBParameterGroupName" : String,
    "DBSecurityGroups" : [ String, ... ],
    "DBSnapshotIdentifier" : String,
    "DBSubnetGroupName" : String,
    "Engine" : String,
    "EngineVersion" : String,
    "Iops" : Number,
    "KmsKeyId" : String,
    "LicenseModel" : String,
    "MasterUsername" : String,
    "MasterUserPassword" : String,
    "MultiAZ" : Boolean,
    "OptionGroupName" : String,
    "Port" : String,
    "PreferredBackupWindow" : String,
    "PreferredMaintenanceWindow" : String,
    "PubliclyAccessible" : Boolean,
    "SourceDBInstanceIdentifier" : String,
    "StorageEncrypted" : Boolean,
    "StorageType" : String,
    "Tags" : [ Resource Tag, ..., ],
    "VPCSecurityGroups" : [ String, ... ]
  }
}
    
```

Properties
----------

 `AllocatedStorage`   
The allocated storage size specified in gigabytes (GB).

If any value is used in the *`Iops`* parameter, *`AllocatedStorage`* must be at least 100 GB, which corresponds to the minimum *`Iops`* value of 1000. If *`Iops`* is increased (in 1000 IOPS increments), then *`AllocatedStorage`* must also be increased (in 100 GB increments) correspondingly.

*Required*: Conditional. This property is required unless you specify the `DBClusterIdentifier` property. In that case, do not specify this property.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AllowMajorVersionUpgrade`   
Indicates whether major version upgrades are allowed. Changing this parameter does not result in an outage, and the change is applied asynchronously as soon as possible.

*Constraints*: This parameter must be set to `true` when you specify an `EngineVersion` that differs from the DB instance's current major version.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AutoMinorVersionUpgrade`   
Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. The default value is `true`.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). For more information, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html) in the *Amazon Relational Database Service API Reference*.

 `AvailabilityZone`   
The name of the Availability Zone where the DB instance is located. You cannot set the *`AvailabilityZone`* parameter if the *`MultiAZ`* parameter is set to true.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `BackupRetentionPeriod`   
The number of days for which automatic DB snapshots are retained.

Important

If this DB instance is deleted or replaced during an update, all automated snapshots are deleted. However, manual DB snapshot are retained.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). For more information, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html) in the *Amazon Relational Database Service API Reference*.

 `CharacterSetName`   
For supported engines, specifies the character set to associate with the database instance. For more information, see [Appendix: Oracle Character Sets Supported in Amazon RDS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html) in the *Amazon Relational Database Service User Guide*.

If you specify the `DBSnapshotIdentifier` or `SourceDBInstanceIdentifier` property, do not specify this property. The value is inherited from the snapshot or source database instance.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBClusterIdentifier`   
The identifier of an existing DB cluster that this instance will be associated with. If you specify this property, specify `aurora` for the `Engine` property and do not specify any of the following properties: `AllocatedStorage`, `CharacterSetName`, `DBSecurityGroups`, `SourceDBInstanceIdentifier`, and `StorageType`.

Amazon RDS assigns the first DB instance in the cluster as the primary and additional DB instances as replicas.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBInstanceClass`   
The name of the compute and memory capacity class of the DB instance.

*Required*: Yes

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `DBInstanceIdentifier`   
A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lower case. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBName`   
The name of the initial database of this instance that was provided at create time, if one was specified. This same name is returned for the life of the DB instance.

Note

If you restore from a snapshot, do specify this property for the MySQL or MariaDB engines.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBParameterGroupName`   
The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](aws-properties-rds-dbparametergroup.html "AWS::RDS::DBParameterGroup") resource created in the template.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). If any of the data members of the referenced parameter group are changed during an update, the database instance might need to be restarted, causing some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.

 `DBSecurityGroups`   
A list of the DB security groups to assign to the Amazon RDS instance. The list can include both the name of existing DB security groups or references to [AWS::RDS::DBSecurityGroup](aws-properties-rds-security-group.html "AWS::RDS::DBSecurityGroup") resources created in the template.

If you set DBSecurityGroups, you must not set [VPCSecurityGroups](aws-properties-rds-database-instance.html#cfn-rds-dbinstance-vpcsecuritygroups), and vice-versa.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `DBSnapshotIdentifier`   
The identifier for the DB snapshot to restore from.

By specifying this property, you can create a DB instance from the specified DB snapshot. If the DBSnapshotIdentifier property is an empty string or the AWS::RDS::DBInstance declaration has no DBSnapshotIdentifier property, the database is created as a new database. If the property contains a value (other than empty string), AWS CloudFormation creates a database from the specified snapshot. If a snapshot with the specified name does not exist, the database creation fails and the stack rolls back.

Some DB instance properties are not valid when you restore from a snapshot, such as the `MasterUsername` and `MasterUserPassword` properties. For information about the properties that you can specify, see the [RestoreDBInstanceFromDBSnapshot](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_RestoreDBInstanceFromDBSnapshot.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBSubnetGroupName`   
A DB subnet group to associate with the DB instance.

If there is no DB subnet group, then it is a non-VPC DB instance.

For more information about using Amazon RDS in a VPC, go to [Using Amazon RDS with Amazon Virtual Private Cloud (VPC)](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon Relational Database Service Developer Guide*.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Engine`   
The name of the database engine that the DB instance uses. This property is optional when you specify the `DBSnapshotIdentifier` property to create DB instances.

For valid values, see the `Engine` parameter of the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: Conditional

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EngineVersion`   
The version number of the database engine to use.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `Iops`   
The number of I/O operations per second (IOPS) that the database provisions. The value must be equal to or greater than `1000`.

If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with `1000` IOPS and `200` GB of storage (a ratio of 5:1) or specify 2000 IOPS with 200 GB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](http://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon Relational Database Service User Guide*.

*Required*: Conditional. If you specify `io1` for the `StorageType` property, you must specify this property.

*Type*: Number

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `KmsKeyId`   
The Amazon Resource Name (ARN) of the AWS Key Management Service master key that is used to encrypt the database instance, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`. If you enable the `StorageEncrypted` property but don't specify this property, the default master key is used. If you specify this property, you must set the `StorageEncrypted` property to `true`.

If you specify the `DBSnapshotIdentifier` or `SourceDBInstanceIdentifier` property, do not specify this property. The value is inherited from the snapshot or source database instance.

Note

Currently, if you specify `DBSecurityGroups`, this property is ignored. If you want to specify a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon Relational Database Service User Guide*.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `LicenseModel`   
The license model information for the DB instance.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `MasterUsername`   
The master user name for the database instance. This property is optional when you specify the `DBSnapshotIdentifier` or the `DBClusterIdentifier` property to create DB instances.

Note

If you specify the `SourceDBInstanceIdentifier` or `DBSnapshotIdentifier` property, do not specify this property. The value is inherited from the source database instance or snapshot.

*Required*: Conditional

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `MasterUserPassword`   
The master password for the database instance. This property is optional when you specify the `DBSnapshotIdentifier` or the `DBClusterIdentifier` property to create DB instances.

Note

If you specify the `SourceDBInstanceIdentifier` property, do not specify this property. The value is inherited from the source database instance.

*Required*: Conditional

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `MultiAZ`   
Specifies if the database instance is a multiple Availability Zone deployment. You cannot set the *`AvailabilityZone`* parameter if the *`MultiAZ`* parameter is set to true.

Note

Do not specify this property if you want a Multi-AZ deployment for a SQL Server database instance. Use the mirroring option in an option group to set Multi-AZ for a SQL Server database instance.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `OptionGroupName`   
An option group that this database instance is associated with.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `Port`   
The port for the instance.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `PreferredBackupWindow`   
The daily time range during which automated backups are created if automated backups are enabled, as determined by the `BackupRetentionPeriod` property. For valid values, see the `PreferredBackupWindow` parameter for the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `PreferredMaintenanceWindow`   
The weekly time range (in UTC) during which system maintenance can occur. For valid values, see the `PreferredMaintenanceWindow` parameter for the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

Note

This property applies during the initial resource creation. If you use AWS CloudFormation to update the DB instance, AWS CloudFormation applies those updates immediately.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). For more information, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html) in the *Amazon Relational Database Service API Reference*.

 `PubliclyAccessible`   
Indicates whether the database instance is an Internet-facing instance. If you specify `true`, an instance is created with a publicly resolvable DNS name, which resolves to a public IP address. If you specify `false`, an internal instance is created with a DNS name that resolves to a private IP address.

The default behavior value depends on your VPC setup and the database subnet group. For more information, see the `PubliclyAccessible` parameter in [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon Relational Database Service API Reference*.

If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the `DependsOn` attribute to declare a dependency on the VPC-gateway attachment. For more information, see [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

Note

Currently, if you specify `DBSecurityGroups`, this property is ignored. If you want to specify a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon Relational Database Service User Guide*.

*Required*: No

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `SourceDBInstanceIdentifier`   
If you want to create a read replica DB instance, specify the ID of the source database instance. Each database instance can have a certain number of read replicas. For more information, see [Working with Read Replicas](http://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon Relational Database Service Developer Guide*.

The `SourceDBInstanceIdentifier` property determines whether a database instance is a read replica. If you remove the `SourceDBInstanceIdentifier` property from your current template and then update your stack, the read replica is deleted and a new database instance (not a read replica) is created.

Important

-   Read replicas do not support deletion policies. Any deletion policy that's associated with a read replica is ignored.

-   If you specify `SourceDBInstanceIdentifier`, do not set the `MultiAZ` property to `true` and do not specify the `DBSnapshotIdentifier` property. You cannot deploy read replicas in multiple Availability Zones, and you cannot create a read replica from a snapshot.

-   Do not set the `BackupRetentionPeriod`, `DBName`, `MasterUsername`, `MasterUserPassword`, and `PreferredBackupWindow` properties. The database attributes are inherited from the source database instance, and backups are disabled for read replicas.

-   If the source DB instance is in a different region than the read replica, specify a valid DB instance ARN. For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon Relational Database Service User Guide*.

-   For DB instances in an Amazon Aurora clusters, do not specify this property. Amazon RDS assigns automatically assigns a writer and reader DB instances.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `StorageEncrypted`   
Indicates whether the database instance is encrypted.

If you specify the `DBClusterIdentifier`, `DBSnapshotIdentifier`, or `SourceDBInstanceIdentifier` property, do not specify this property. The value is inherited from the cluster, snapshot, or source database instance.

*Required*: Conditional. If you specify the `KmsKeyId` property, you must enable encryption.

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `StorageType`   
The storage type associated with this database instance.

For the default and valid values, see the `StorageType` parameter of the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this database instance.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `VPCSecurityGroups`   
A list of the VPC security group IDs to assign to the Amazon RDS instance. The list can include both the physical IDs of existing VPC security groups or references to [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup") resources created in the template.

If you set VPCSecurityGroups, you must not set [DBSecurityGroups](aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups), and vice-versa.

Important

You can migrate a database instance in your stack from an RDS DB security group to a VPC security group, but you should keep the following points in mind:

-   You cannot revert to using an RDS security group once you have established a VPC security group membership.

-   When you migrate your DB instance to VPC security groups, if your stack update rolls back because of another failure in the database instance update, or because of an update failure in another AWS CloudFormation resource, the rollback will fail because it cannot revert to an RDS security group.

To avoid this situation, only migrate your DB instance to using VPC security groups when that is the *only* change in your stack template.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Updating and Deleting AWS::RDS::DBInstance resources
----------------------------------------------------

When updates are made to properties labeled "*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)", AWS CloudFormation first creates a replacement DB instance resource, then changes references from other dependent resources to point to the replacement resource, and finally deletes the old resource.

Caution

If you do not take a snapshot of the database before updating the stack, you will lose the data when your DB instance is replaced. To preserve your data, take the following precautions:

1.  Deactivate any applications that are using the DB instance so that there is no activity against the DB instance.

2.  Create a snapshot of the DB instance. For more information about creating DB snapshots, see [Creating a DB snapshot](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html).

3.  If you want to restore your instance using a DB snapshot, modify the update template with your DB instance changes and add the DBSnapshotIdentifier property with the ID of the DB snapshot that you want to use.

4.  Update the stack.

For more information about updating other properties on this resource, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html). For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

You can set a deletion policy for your DB instance to control how AWS CloudFormation handles the instance when the stack is deleted. For Amazon RDS DB instances, you can choose to *retain* the instance, to *delete* the instance, or to *create a snapshot* of the instance. For more information, see [DeletionPolicy Attribute](aws-attribute-deletionpolicy.html "DeletionPolicy Attribute").

Return Values
-------------

### Ref

When you provide the RDS DB instance's logical name to the `Ref` intrinsic function, `Ref` will return the DBInstanceIdentifier. For example: `mystack-mydb-ea5ugmfvuaxg`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

-   **Endpoint.Address**

    The connection endpoint for the database. For example: `mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-1.rds.amazonaws.com`.

-   **Endpoint.Port**

    The port number on which the database accepts connections. For example: `3306`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

**Example DBInstance with a set MySQL version, Tags and DeletionPolicy**

This example shows how to set the MySQL version that has a `DeletionPolicy Attribute` set. With the `DeletionPolicy` set to `Snapshot`, AWS CloudFormation will take a snapshot of this DB instance before deleting it during stack deletion. A tag that contains a friendly name for the database is also set.

``` {.programlisting}
          
"MyDB" : {
   "Type" : "AWS::RDS::DBInstance",
   "Properties" : {
      "DBName" : { "Ref" : "DBName" },
      "AllocatedStorage" : { "Ref" : "DBAllocatedStorage" },
      "DBInstanceClass" : { "Ref" : "DBInstanceClass" },
      "Engine" : "MySQL",
      "EngineVersion" : "5.5",
      "MasterUsername" : { "Ref" : "DBUser" },
      "MasterUserPassword" : { "Ref" : "DBPassword" },
      "Tags" : [ { "Key" : "Name", "Value" : "My SQL Database" } ]
   },
   "DeletionPolicy" : "Snapshot"
}         
        
```

**Example DBInstance with provisioned IOPS**

This example sets a provisioned IOPS value in the [Iops](aws-properties-rds-database-instance.html#cfn-rds-dbinstance-iops) property. Note that the [AllocatedStorage](aws-properties-rds-database-instance.html#cfn-rds-dbinstance-allocatedstorage) property is set according to the 10:1 ratio between IOPS and GiBs of storage.

``` {.programlisting}
          
"MyDB" : {
   "Type" : "AWS::RDS::DBInstance",
   "Properties" : {
      "AllocatedStorage" : "100",
      "DBInstanceClass" : "db.m1.small",
      "Engine" : "MySQL",
      "EngineVersion" : "5.5",
      "Iops" : "1000",
      "MasterUsername" : { "Ref" : "DBUser" },
      "MasterUserPassword" : { "Ref" : "DBPassword" }
   }
}        
        
```

**Example Read replica DBInstance**

This example creates a read replica named `MyDBreadreplica` for the `MyDB` DB instance.

``` {.programlisting}
          "MyDB" : {
   "Type" : "AWS::RDS::DBInstance",
   "Properties" : {
      "DBName" : { "Ref" : "DBName" },
      "AllocatedStorage" : { "Ref" : "DBAllocatedStorage" },
      "DBInstanceClass" : { "Ref" : "DBClass" },
      "Engine" : "MySQL",
      "EngineVersion" : "5.6",
      "MasterUsername" : { "Ref" : "DBUser" } ,
      "MasterUserPassword" : { "Ref" : "DBPassword" },
      "Port" : "5804",
      "Tags" : [{"Key" : "Role", "Value" : "Primary"}] 
   }
},

"MyDBreadreplica" : {
   "Type": "AWS::RDS::DBInstance",
   "Properties": {
      "SourceDBInstanceIdentifier": { "Ref" : "MyDB" },
      "Port" : "5802",
      "Tags" : [{"Key" : "Role", "Value" : "ReadRep"}]     
      }
   }
}
        
```

To view more AWS::RDS::DBInstance template snippets, see [Amazon RDS Template Snippets](quickref-rds.html "Amazon RDS Template Snippets").

