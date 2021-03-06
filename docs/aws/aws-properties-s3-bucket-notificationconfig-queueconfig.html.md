Amazon Simple Storage Service NotificationConfiguration QueueConfigurations
===========================================================================

`QueueConfigurations` is a property of the [Amazon S3 NotificationConfiguration](aws-properties-s3-bucket-notificationconfig.html "Amazon S3 NotificationConfiguration") property that describes the S3 bucket events about which you want to send messages to Amazon SQS and the queues to which you want to send them.

Syntax
------

``` {.programlisting}
      {
  "Event" : String,
  "Filter" : Filter,
  "Queue" : String 
}
    
```

Properties
----------

 `Event`   
The S3 bucket event about which you want to publish messages to Amazon Simple Queue Service ( Amazon SQS). For more information, see [Supported Event Types](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: Yes

*Type*: String

 `Filter`   
The filtering rules that determine for which objects to send notifications. For example, you can create a filter so that Amazon Simple Storage Service (Amazon S3) sends notifications only when image files with a `.jpg` extension are added to the bucket.

*Required*: No

*Type*: [Amazon S3 NotificationConfiguration Config Filter](aws-properties-s3-bucket-notificationconfiguration-config-filter.html "Amazon S3 NotificationConfiguration Config Filter")

 `Queue`   
The Amazon Resource Name (ARN) of the Amazon SQS queue that Amazon S3 publishes messages to when the specified event type occurs.

*Required*: Yes

*Type*: String


