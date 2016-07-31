Amazon S3 NotificationConfiguration
===================================

Describes the notification configuration for an [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource.

Syntax
------

``` {.programlisting}
      {
  "LambdaConfigurations" : [ Lambda Configuration, ... ],
  "QueueConfigurations" : [ Queue Configuration, ... ],
  "TopicConfigurations" : [ Topic Configuration, ... ]
}
    
```

Properties
----------

 `LambdaConfigurations`   
The AWS Lambda functions to invoke and the events for which to invoke the functions.

*Required*: No

*Type*: [Amazon Simple Storage Service NotificationConfiguration LambdaConfigurations](aws-properties-s3-bucket-notificationconfig-lambdaconfig.html "Amazon Simple Storage Service NotificationConfiguration LambdaConfigurations")

 `QueueConfigurations`   
The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages.

*Required*: No

*Type*: [Amazon Simple Storage Service NotificationConfiguration QueueConfigurations](aws-properties-s3-bucket-notificationconfig-queueconfig.html "Amazon Simple Storage Service NotificationConfiguration QueueConfigurations")

 `TopicConfigurations`   
The topic to which notifications are sent and the events for which notification are generated.

*Required*: No

*Type*: [Amazon S3 NotificationConfiguration TopicConfigurations](aws-properties-s3-bucket-notificationconfig-topicconfig.html "Amazon S3 NotificationConfiguration TopicConfigurations")


