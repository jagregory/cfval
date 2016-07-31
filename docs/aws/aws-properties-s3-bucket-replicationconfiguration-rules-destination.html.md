Amazon S3 ReplicationConfiguration Rules Destination
====================================================

`Destination` is a property of the [Amazon S3 ReplicationConfiguration Rules](aws-properties-s3-bucket-replicationconfiguration-rules.html "Amazon S3 ReplicationConfiguration Rules") property that specifies which Amazon Simple Storage Service (Amazon S3) bucket to store replicated objects and their storage class.

Syntax
------

``` {.programlisting}
      {
  "Bucket" : String,
  "StorageClass" : String
}
    
```

Properties
----------

 `Bucket`   
The Amazon resource name (ARN) of an S3 bucket where Amazon S3 stores replicated objects. This destination bucket must be in a different region than your source bucket.

If you have multiple rules in your replication configuration, specify the same destination bucket for all of the rules.

*Required*: Yes

*Type*: String

 `StorageClass`   
The storage class to use when replicating objects, such as standard or reduced redundancy. By default, Amazon S3 uses the storage class of the source object to create object replica. For valid values, see the `StorageClass` element of the [PUT Bucket replication](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon Simple Storage Service API Reference*.

*Required*: No

*Type*: String


