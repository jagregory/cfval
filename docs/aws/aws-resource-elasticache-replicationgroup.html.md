AWS::ElastiCache::ReplicationGroup
==================================

The `AWS::ElastiCache::ReplicationGroup` resource creates an Amazon ElastiCache replication group. A replication group is a collection of cache clusters, where one of the clusters is a primary read-write cluster and the others are read-only replicas.

Note

Currently, replication groups are supported only for Redis clusters.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ElastiCache::ReplicationGroup",
  "Properties" : {
    "AutomaticFailoverEnabled" : Boolean,
    "AutoMinorVersionUpgrade" : Boolean,
    "CacheNodeType" : String,
    "CacheParameterGroupName" : String,
    "CacheSecurityGroupNames" : [ String, ... ],
    "CacheSubnetGroupName" : String,
    "Engine" : String,
    "EngineVersion" : String,
    "NotificationTopicArn" : String,
    "NumCacheClusters" : Integer,
    "Port" : Integer,
    "PreferredCacheClusterAZs" : [ String, ... ],
    "PreferredMaintenanceWindow" : String,
    "ReplicationGroupDescription" : String,
    "SecurityGroupIds" : [ String, ... ],
    "SnapshotArns" : [ String, ... ],
    "SnapshotRetentionLimit" : Integer,
    "SnapshotWindow" : String
  }
}
    
```

Properties
----------

For more information about each property and valid values, see [CreateReplicationGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateReplicationGroup.html) in the *Amazon ElastiCache API Reference Guide*.

 `AutomaticFailoverEnabled`   
Indicates whether Multi-AZ is enabled. When Multi-AZ is enabled, a read-only replica is automatically promoted to a read-write primary cluster if the existing primary cluster fails. If you specify `true`, you must specify a value greater than `1` for the `NumCacheNodes` property. By default, AWS CloudFormation sets the value to `true`.

For more information about Multi-AZ, see [Multi-AZ with Redis Replication Groups](http://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/AutoFailover.html) in the *Amazon ElastiCache User Guide*.

Note

You cannot enable automatic failover for Redis versions earlier than 2.8.6 or for T1 and T2 cache node types.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AutoMinorVersionUpgrade`   
Currently, this property isn't used by ElastiCache.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `CacheNodeType`   
The compute and memory capacity of nodes in the node group. To see valid values, see [CreateReplicationGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateReplicationGroup.html) in the *Amazon ElastiCache API Reference Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `CacheParameterGroupName`   
The name of the parameter group to associate with this replication group.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `CacheSecurityGroupNames`   
A list of cache security group names to associate with this replication group. If you specify the `SecurityGroupIds` property, do not specify this property; you can specify only one.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `CacheSubnetGroupName`   
The name of a cache subnet group to use for this replication group.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Engine`   
The name of the cache engine to use for the cache clusters in this replication group. Currently, you can specify only `redis`.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EngineVersion`   
The version number of the cache engine to use for the cache clusters in this replication group.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NotificationTopicArn`   
The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic to which notifications are sent.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NumCacheClusters`   
The number of cache clusters for this replication group. If automatic failover is enabled, you must specify a value greater than `1`.

*Required*: Yes

*Type*: Integer

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Port`   
The port number on which each member of the replication group accepts connections.

*Required*: No

*Type*: Integer

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `PreferredCacheClusterAZs`   
A list of Availability Zones (AZs) in which the cache clusters in this replication group are created.

*Required*: No

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `PreferredMaintenanceWindow`   
The weekly time range during which system maintenance can occur. Use the following format to specify a time range: `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). For example, you can specify `sun:22:00-sun:23:30` for Sunday from 10 PM to 11:30 PM.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `ReplicationGroupDescription`   
The description of the replication group.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SecurityGroupIds`   
A list of Amazon Virtual Private Cloud (Amazon VPC) security groups to associate with this replication group. Use this property only when you are creating a replication group in a VPC. If you specify the `CacheSecurityGroupNames` property, do not specify this property; you can specify only one.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SnapshotArns`   
A single-element string list that specifies an ARN of a Redis `.rdb` snapshot file that is stored in Amazon Simple Storage Service (Amazon S3). The snapshot file populates the node group. The Amazon S3 object name in the ARN cannot contain commas. For example, you can specify `arn:aws:s3:::my_bucket/snapshot1.rdb`.

*Required*: No

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SnapshotRetentionLimit`   
The number of days that ElastiCache retains automatic snapshots before deleting them.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SnapshotWindow`   
The time range (in UTC) when ElastiCache takes a daily snapshot of your node group. For example, you can specify `05:00-09:00`.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

In the following sample, the `Ref` function returns the name of the `myReplicationGroup` replication group, such as `abc12xmy3d1w3hv6`.

``` {.programlisting}
        { "Ref": "myReplicationGroup" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `PrimaryEndPoint.Address`   
The DNS address of the primary read-write cache node.

 `PrimaryEndPoint.Port`   
The number of the port that the primary read-write cache engine is listening on.

 `ReadEndPoint.Addresses`   
A string with a list of endpoints for the read-only replicas. The order of the addresses map to the order of the ports from the `ReadEndPoint.Ports` attribute.

 `ReadEndPoint.Ports`   
A string with a list of ports for the read-only replicas. The order of the ports map to the order of the addresses from the `ReadEndPoint.Addresses` attribute.

 `ReadEndPoint.Addresses.List`   
A list of endpoints for the read-only replicas. The order of the addresses map to the order of the ports from the `ReadEndPoint.Ports.List` attribute.

 `ReadEndPoint.Ports.List`   
A list of ports for the read-only replicas. The order of the ports map to the order of the addresses from the `ReadEndPoint.Addresses.List` attribute.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Example
-------

The following sample declares a replication group with two nodes and automatic failover enabled.

``` {.programlisting}
      "myReplicationGroup" : {
  "Type": "AWS::ElastiCache::ReplicationGroup",
  "Properties": {
    "ReplicationGroupDescription" : "my description",
    "NumCacheClusters" : "2",
    "Engine" : "redis",
    "CacheNodeType" : "cache.m3.medium",
    "AutoMinorVersionUpgrade" : "true",
    "AutomaticFailoverEnabled" : "true",
    "CacheSubnetGroupName" : "subnetgroup",
    "EngineVersion" : "2.8.6",
    "PreferredMaintenanceWindow" : "wed:09:25-wed:22:30",
    "SnapshotRetentionLimit" : "4",
    "SnapshotWindow" : "03:30-05:30"
  }
}
    
```
