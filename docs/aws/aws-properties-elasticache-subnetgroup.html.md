AWS::ElastiCache::SubnetGroup
=============================

Creates a cache subnet group. For more information about cache subnet groups, go to [Cache Subnet Groups](http://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/CacheSubnetGroups.html) in the *Amazon ElastiCache User Guide* or go to [CreateCacheSubnetGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html) in the *Amazon ElastiCache API Reference Guide*.

When you specify an AWS::ElastiCache::SubnetGroup type as an argument to the `Ref` function, AWS CloudFormation returns the name of the cache subnet group.

Syntax
------

``` {.programlisting}
      "SubnetGroup" : {
    "Type" : "AWS::ElastiCache::SubnetGroup",
    "Properties" : {
        "Description" : String,
        "SubnetIds" : [ String, ... ]
    }
}
    
```

Properties
----------

 `Description`   
The description for the cache subnet group.

*Type*: String

*Required*: Yes

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SubnetIds`   
The Amazon EC2 subnet IDs for the cache subnet group.

*Type*: String list

*Required*: Yes

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Example
-------

``` {.programlisting}
      "SubnetGroup" : {
    "Type" : "AWS::ElastiCache::SubnetGroup",
    "Properties" : {
        "Description" : "Cache Subnet Group",
        "SubnetIds" : [ { "Ref" : "Subnet1" }, { "Ref" : "Subnet2" } ]
    }
}
    
```
