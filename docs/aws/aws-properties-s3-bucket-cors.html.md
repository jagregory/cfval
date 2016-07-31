Amazon S3 Cors Configuration
============================

Describes the cross-origin access configuration for objects in an [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource.

Syntax
------

``` {.programlisting}
      {
  "CorsRules" : [ CorsRules, ... ]
}
    
```

Properties
----------

 `CorsRules`   
A set of origins and methods that you allow.

*Required*: Yes

*Type*: [Amazon S3 Cors Configuration Rule](aws-properties-s3-bucket-cors-corsrule.html "Amazon S3 Cors Configuration Rule")


