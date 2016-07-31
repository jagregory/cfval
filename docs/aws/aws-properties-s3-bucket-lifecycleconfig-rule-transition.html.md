Amazon S3 Lifecycle Rule Transition
===================================

Describes when an object transitions to a specified storage class for the [Amazon S3 Lifecycle Rule](aws-properties-s3-bucket-lifecycleconfig-rule.html "Amazon S3 Lifecycle Rule") property.

Syntax
------

``` {.programlisting}
      {
  "StorageClass" : String,
  "TransitionDate" : String,
  "TransitionInDays" : Integer
}
    
```

Properties
----------

 `StorageClass`   
The storage class to which you want the object to transition, such as `GLACIER`. For valid values, see the `StorageClass` request element of the [PUT Bucket lifecycle](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) action in the *Amazon Simple Storage Service API Reference*.

*Required*: Yes

*Type*: String

 `TransitionDate`   
Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format. The time is always midnight UTC.

*Required*: Conditional

*Type*: String

 `TransitionInDays`   
Indicates the number of days after creation when objects are transitioned to the specified storage class.

*Required*: Conditional

*Type*: Integer


