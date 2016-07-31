Amazon S3 ReplicationConfiguration Rules
========================================

`Rules` is a property of the [Amazon S3 ReplicationConfiguration](aws-properties-s3-bucket-replicationconfiguration.html "Amazon S3 ReplicationConfiguration") property that specifies which Amazon Simple Storage Service (Amazon S3) objects to replicate and where to store them.

Syntax
------

``` {.programlisting}
      {
  "Destination" : String,
  "Id" : String,
  "Prefix" : String,
  "Status" : String
}
    
```

Properties
----------

 `Destination`   
Defines the destination where Amazon S3 stores replicated objects.

*Required*: Yes

*Type*: [Amazon S3 ReplicationConfiguration Rules Destination](aws-properties-s3-bucket-replicationconfiguration-rules-destination.html "Amazon S3 ReplicationConfiguration Rules Destination")

 `Id`   
A unique identifier for the rule. If you don't specify a value, AWS CloudFormation generates a random ID.

*Required*: No

*Type*: String

 `Prefix`   
An object prefix. This rule applies to all Amazon S3 objects with this prefix. To specify all objects in an S3 bucket, specify an empty string.

*Required*: Yes

*Type*: String

 `Status`   
Whether the rule is enabled. For valid values, see the `Status` element of the [PUT Bucket replication](http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon Simple Storage Service API Reference*.

*Required*: Yes

*Type*: String


