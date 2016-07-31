AWS::ElastiCache::ParameterGroup
================================

The AWS::ElastiCache::ParameterGroup type creates a new cache parameter group. Cache parameter groups control the parameters for a cache cluster.

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::ElastiCache::ParameterGroup",
   "Properties": {
      "CacheParameterGroupFamily" : String,
      "Description" : String,
      "Properties" : { String:String, ... }
   }
}     
    
```

Properties
----------

 `CacheParameterGroupFamily`   
The name of the cache parameter group family that the cache parameter group can be used with.

*Required*: Yes

*Type*: String

*Update requires*: Updates are not supported.

 `Description`   
The description for the Cache Parameter Group.

*Required*: Yes

*Type*: String

*Update requires*: Updates are not supported.

 `Properties`   
A comma-delimited list of parameter name/value pairs. For more information, go to [ModifyCacheParameterGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html) in the *Amazon ElastiCache API Reference Guide*.

*Example*:

``` {.programlisting}
            
"Properties" : {
   "cas_disabled" : "1",
   "chunk_size_growth_factor" : "1.02"
}              
          
```

*Required*: No

*Type*: Mapping of key-value pairs

*Update requires*: Updates are not supported.

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      "MyParameterGroup": {
   "Type": "AWS::ElastiCache::ParameterGroup",
   "Properties": {
      "Description": "MyNewParameterGroup",
      "CacheParameterGroupFamily": "memcached1.4",
      "Properties" : {
         "cas_disabled" : "1",
         "chunk_size_growth_factor" : "1.02"
      }
   }
}
    
```

See Also
--------

-   [CreateCacheParameterGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheParameterGroup.html) in the *Amazon ElastiCache API Reference Guide*

-   [ModifyCacheParameterGroup](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html) in the *Amazon ElastiCache API Reference Guide*

-   [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates")


