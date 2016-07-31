Amazon S3 Lifecycle Rule NoncurrentVersionTransition
====================================================

`NoncurrentVersionTransition` is a property of the [Amazon S3 Lifecycle Rule](aws-properties-s3-bucket-lifecycleconfig-rule.html "Amazon S3 Lifecycle Rule") property that describes when noncurrent objects transition to a specified storage class.

Syntax
------

``` {.programlisting}
      {
  "StorageClass" : String,
  "TransitionInDays" : Integer
}
    
```

Properties
----------

 `StorageClass`   
The storage class to which you want the object to transition, such as `GLACIER`. For valid values, see the `StorageClass` request element of the [PUT Bucket lifecycle](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) action in the *Amazon Simple Storage Service API Reference*.

*Required*: Yes

*Type*: String

 `TransitionInDays`   
The number of days between the time that a new version of the object is uploaded to the bucket and when old versions of the object are transitioned to the specified storage class.

*Required*: Yes

*Type*: Integer


