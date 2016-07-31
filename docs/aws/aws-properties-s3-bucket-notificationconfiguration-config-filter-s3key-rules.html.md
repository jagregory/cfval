Amazon S3 NotificationConfiguration Config Filter S3Key Rules
=============================================================

`Rules` is a property of the [Amazon S3 NotificationConfiguration Config Filter S3Key](aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html "Amazon S3 NotificationConfiguration Config Filter S3Key") property that describes the Amazon Simple Storage Service (Amazon S3) object key name to filter on and whether to filter on the suffix or prefix of the key name.

Syntax
------

``` {.programlisting}
      {
  "Name" : String,
  "Value" : String
}
    
```

Properties
----------

 `Name`   
Whether the filter matches the prefix or suffix of object key names. For valid values, see the `Name` request element of the [PUT Bucket notification](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTnotification.html) action in the *Amazon Simple Storage Service API Reference*.

*Required*: Yes

*Type*: String

 `Value`   
The value that the filter searches for in object key names.

*Required*: Yes

*Type*: String


