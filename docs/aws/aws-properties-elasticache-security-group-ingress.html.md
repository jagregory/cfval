AWS::ElastiCache::SecurityGroupIngress
======================================

The AWS::ElastiCache::SecurityGroupIngress type authorizes ingress to a cache security group from hosts in specified Amazon EC2 security groups. For more information about ElastiCache security group ingress, go to [AuthorizeCacheSecurityGroupIngress](http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_AuthorizeCacheSecurityGroupIngress.html) in the *Amazon ElastiCache API Reference Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ElastiCache::SecurityGroupIngress",
  "Properties" :
  {
    "CacheSecurityGroupName" : String,
    "EC2SecurityGroupName" : String,
    "EC2SecurityGroupOwnerId" : String
  }
}
    
```

Properties
----------

 `CacheSecurityGroupName`   
The name of the Cache Security Group to authorize.

*Type*: String

*Required*: Yes

*Update requires*: Updates are not supported.

 `EC2SecurityGroupName`   
Name of the EC2 Security Group to include in the authorization.

*Type*: String

*Required*: Yes

*Update requires*: Updates are not supported.

 `EC2SecurityGroupOwnerId`   
Specifies the AWS Account ID of the owner of the EC2 security group specified in the EC2SecurityGroupName property. The AWS access key ID is not an acceptable value.

*Type*: String

*Required*: No

*Update requires*: Updates are not supported.


