AWS::ElastiCache::SecurityGroup
===============================

The `AWS::ElastiCache::SecurityGroup` resource creates a cache security group. For more information about cache security groups, go to [Cache Security Groups](http://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/CacheSecurityGroup.html) in the *Amazon ElastiCache User Guide* or go to [CreateCacheSecurityGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSecurityGroup.html) in the *Amazon ElastiCache API Reference Guide*.

To create an ElastiCache cluster in a VPC, use the [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup") resource. For more information, see the `VpcSecurityGroupIds` property in the [AWS::ElastiCache::CacheCluster](aws-properties-elasticache-cache-cluster.html "AWS::ElastiCache::CacheCluster") resource.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ElastiCache::SecurityGroup",
  "Properties" :
  {
    "Description" : String
  }
}
    
```

Properties
----------

 `Description`   
A description for the cache security group.

*Type*: String

*Required*: No

*Update requires*: Updates are not supported.

Return Values
-------------

### Ref

When you specify the `AWS::ElastiCache::SecurityGroup` resource as an argument to the `Ref` function, AWS CloudFormation returns the *`CacheSecurityGroupName`* property of the cache security group.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

