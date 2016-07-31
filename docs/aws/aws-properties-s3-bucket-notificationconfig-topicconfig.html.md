Amazon S3 NotificationConfiguration TopicConfigurations
=======================================================

Describes the topic and events for the [Amazon S3 NotificationConfiguration](aws-properties-s3-bucket-notificationconfig.html "Amazon S3 NotificationConfiguration") property.

Syntax
------

``` {.programlisting}
      {
  "Event" : String,
  "Filter" : Filter,
  "Topic" : String 
}
    
```

Properties
----------

 `Event`   
The Amazon Simple Storage Service (Amazon S3) bucket event about which to send notifications. For more information, see [Supported Event Types](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: Yes

*Type*: String

 `Filter`   
The filtering rules that determine for which objects to send notifications. For example, you can create a filter so that Amazon Simple Storage Service (Amazon S3) sends notifications only when image files with a `.jpg` extension are added to the bucket.

*Required*: No

*Type*: [Amazon S3 NotificationConfiguration Config Filter](aws-properties-s3-bucket-notificationconfiguration-config-filter.html "Amazon S3 NotificationConfiguration Config Filter")

 `Topic`   
The Amazon SNS topic Amazon Resource Name (ARN) to which Amazon S3 reports the specified events.

*Required*: Yes

*Type*: String


