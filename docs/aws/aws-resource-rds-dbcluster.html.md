AWS::RDS::DBCluster
===================

The `AWS::RDS::DBCluster` resource creates a cluster, such as an Aurora for Amazon RDS (Amazon Aurora) DB cluster. Amazon Aurora is a fully managed, MySQL-compatible, relational database engine. For more information, see [Aurora on Amazon RDS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html) in the *Amazon Relational Database Service User Guide*.

Note

Currently, you can create this resource only in regions in which Amazon Aurora is supported.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::RDS::DBCluster",
  "Properties" :
  {
    "AvailabilityZones" : [ String, ... ],
    "BackupRetentionPeriod" : Integer,
    "DatabaseName" : String,
    "DBClusterParameterGroupName" : String,
    "DBSubnetGroupName" : String,
    "Engine" : String,
    "EngineVersion" : String,
    "KmsKeyId" : String,
    "MasterUsername" : String,
    "MasterUserPassword" : String,
    "Port" : Integer,
    "PreferredBackupWindow" : String,
    "PreferredMaintenanceWindow" : String,
    "SnapshotIdentifier" : String,
    "StorageEncrypted" : Boolean,
    "Tags" : [ Resource Tag, ... ],
    "VpcSecurityGroupIds" : [ String, ... ]
  }
}
    
```

Properties
----------

 `AvailabilityZones`   
A list of Availability Zones (AZs) in which DB instances in the cluster can be created.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `BackupRetentionPeriod`   
The number of days for which automatic backups are retained. For more information, see [CreateDBCluster](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html) in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). For more information, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html) in the *Amazon Relational Database Service API Reference*.

 `DatabaseName`   
The name of your database. You can specify a name of up to eight alpha-numeric characters. If you do not provide a name, Amazon Relational Database Service (Amazon RDS) won't create a database in this DB cluster.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBClusterParameterGroupName`   
The name of the DB cluster parameter group to associate with this DB cluster. For the default value, see the `DBClusterParameterGroupName` parameter of the [CreateDBCluster](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `DBSubnetGroupName`   
A DB subnet group that you want to associate with this DB cluster.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Engine`   
The name of the database engine that you want to use for this DB cluster.

For valid values, see the `Engine` parameter of the [CreateDBCluster](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EngineVersion`   
The version number of the database engine that you want to use.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `KmsKeyId`   
The Amazon Resource Name (ARN) of the AWS Key Management Service master key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`. If you enable the `StorageEncrypted` property but don't specify this property, the default master key is used. If you specify this property, you must set the `StorageEncrypted` property to `true`.

If you specify the `SnapshotIdentifier`, do not specify this property. The value is inherited from the snapshot DB cluster.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `MasterUsername`   
The master user name for the DB instance.

*Required*: Conditional. You must specify this property unless you specify the `SnapshotIdentifier` property. In that case, do not specify this property.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `MasterUserPassword`   
The password for the master database user.

*Required*: Conditional. You must specify this property unless you specify the `SnapshotIdentifier` property. In that case, do not specify this property.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Port`   
The port number on which the DB instances in the cluster can accept connections.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PreferredBackupWindow`   
if automated backups are enabled (see the `BackupRetentionPeriod` property), the daily time range in UTC during which you want to create automated backups.

For valid values, see the `PreferredBackupWindow` parameter of the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PreferredMaintenanceWindow`   
The weekly time range (in UTC) during which system maintenance can occur.

For valid values, see the `PreferredMaintenanceWindow` parameter of the [CreateDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) action in the *Amazon Relational Database Service API Reference*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). For more information, see [ModifyDBInstance](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBInstance.html) in the *Amazon Relational Database Service API Reference*.

 `SnapshotIdentifier`   
The identifier for the DB cluster snapshot from which you want to restore.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `StorageEncrypted`   
Indicates whether the DB instances in the cluster are encrypted.

If you specify the `SnapshotIdentifier` property, do not specify this property. The value is inherited from the snapshot DB cluster.

*Required*: Conditional. If you specify the `KmsKeyId` property, you must enable encryption.

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `Tags`   
The tags that you want to attach to this DB cluster.

*Required*: No

*Type*: A list of [resource tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: Updates are not supported.

 `VpcSecurityGroupIds`   
A list of VPC security groups to associate with this DB cluster.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

-   **Endpoint.Address**

    The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-1.rds.amazonaws.com`.

-   **Endpoint.Port**

    The number of the port on which the DB cluster accepts connections, such as `3306`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Example
-------

The following snippet creates an Amazon Aurora DB cluster and adds two DB instances to it. Because Amazon RDS automatically assigns a writer and reader DB instances in the cluster, use the cluster endpoint to read and write data, not the individual DB instance endpoints.

``` {.programlisting}
      "RDSCluster" : {
  "Type" : "AWS::RDS::DBCluster",
  "Properties" : {
    "MasterUsername" : { "Ref" : "username" },
    "MasterUserPassword" : { "Ref" : "password" },
    "Engine" : "aurora",
    "DBSubnetGroupName" : { "Ref" : "DBSubnetGroup" },
    "DBClusterParameterGroupName" : { "Ref" : "RDSDBClusterParameterGroup" }
  }
},
"RDSDBInstance1" : {
  "Type" : "AWS::RDS::DBInstance",
  "Properties" : {
    "DBSubnetGroupName" : {
      "Ref" : "DBSubnetGroup"
    },
    "Engine" : "aurora",
    "DBClusterIdentifier" : {
      "Ref" : "RDSCluster"
    },
    "PubliclyAccessible" : "true",
    "AvailabilityZone" : { "Fn::GetAtt" : [ "Subnet1", "AvailabilityZone" ] },
    "DBInstanceClass" : "db.r3.xlarge"
  }
},
"RDSDBInstance2" : {
  "Type" : "AWS::RDS::DBInstance",
  "Properties" : {
    "DBSubnetGroupName" : {
      "Ref" : "DBSubnetGroup"
    },
    "Engine" : "aurora",
    "DBClusterIdentifier" : {
      "Ref" : "RDSCluster"
    },
    "PubliclyAccessible" : "true",
    "AvailabilityZone" : { "Fn::GetAtt" : [ "Subnet2", "AvailabilityZone" ] },
    "DBInstanceClass" : "db.r3.xlarge"
  }
}

    
```
