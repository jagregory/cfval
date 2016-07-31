Amazon S3 NotificationConfiguration Config Filter
=================================================

`Filter` is a property of the `LambdaConfigurations`, `QueueConfigurations`, and `TopicConfigurations` properties that describes the filtering rules that determine the Amazon Simple Storage Service (Amazon S3) objects for which to send notifications.

Syntax
------

``` {.programlisting}
      {
  "S3Key" : S3 Key
}
    
```

Properties
----------

 `S3Key`   
Amazon S3 filtering rules that describe for which object key names to send notifications.

*Required*: Yes

*Type*: [Amazon S3 NotificationConfiguration Config Filter S3Key](aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html "Amazon S3 NotificationConfiguration Config Filter S3Key")


