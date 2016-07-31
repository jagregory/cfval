Amazon S3 Versioning Configuration
==================================

Describes the versioning state of an [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource. For more information, see [PUT Bucket versioning](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the *Amazon Simple Storage Service API Reference*.

Syntax
------

``` {.programlisting}
      {
  "Status" : String
}
    
```

Properties
----------

 `Status`   
The versioning state of an Amazon S3 bucket. If you enable versioning, you must suspend versioning to disable it.

*Required*: Yes

*Type*: String


