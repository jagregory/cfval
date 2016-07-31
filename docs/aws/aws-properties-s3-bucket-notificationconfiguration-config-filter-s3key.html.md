Amazon S3 NotificationConfiguration Config Filter S3Key
=======================================================

`S3Key` is a property of the [Amazon S3 NotificationConfiguration Config Filter](aws-properties-s3-bucket-notificationconfiguration-config-filter.html "Amazon S3 NotificationConfiguration Config Filter") property that specifies the key names of Amazon Simple Storage Service (Amazon S3) objects for which to send notifications.

Syntax
------

``` {.programlisting}
      {
  "Rules" : [ Rule, ... ]
}
    
```

Properties
----------

 `Rules`   
The object key name to filter on and whether to filter on the suffix or prefix of the key name.

*Required*: Yes

*Type*: List of [Amazon S3 NotificationConfiguration Config Filter S3Key Rules](aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html "Amazon S3 NotificationConfiguration Config Filter S3Key Rules")


