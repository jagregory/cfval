Amazon S3 Lifecycle Configuration
=================================

Describes the lifecycle configuration for objects in an [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource.

Syntax
------

``` {.programlisting}
      {
  "Rules" : [ Lifecycle Rule, ... ]
}
    
```

Properties
----------

 `Rules`   
A lifecycle rule for individual objects in an S3 bucket.

*Required*: Yes

*Type*: [Amazon S3 Lifecycle Rule](aws-properties-s3-bucket-lifecycleconfig-rule.html "Amazon S3 Lifecycle Rule")


