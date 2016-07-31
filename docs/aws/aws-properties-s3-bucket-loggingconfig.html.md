Amazon S3 Logging Configuration
===============================

Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for an [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource. These logs track requests to an Amazon S3 bucket. For more information, see [PUT Bucket logging](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the *Amazon Simple Storage Service API Reference*.

Syntax
------

``` {.programlisting}
      {
  "DestinationBucketName" : String,
  "LogFilePrefix" : String
}
    
```

Properties
----------

 `DestinationBucketName`   
The name of an Amazon S3 bucket where Amazon S3 store server access log files. You can store log files in any bucket that you own. By default, logs are stored in the bucket where the `LoggingConfiguration` property is defined.

*Required*: No

*Type*: String

 `LogFilePrefix`   
A prefix for the all log object keys. If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.

*Required*: No

*Type*: String


