Amazon Simple Storage Service NotificationConfiguration LambdaConfigurations
============================================================================

`LambdaConfigurations` is a property of the [Amazon S3 NotificationConfiguration](aws-properties-s3-bucket-notificationconfig.html "Amazon S3 NotificationConfiguration") property that describes the AWS Lambda (Lambda) functions to invoke and the events for which to invoke them.

Syntax
------

``` {.programlisting}
      {
  "Event" : String,
  "Filter" : Filter,
  "Function" : String 
}
    
```

Properties
----------

 `Event`   
The S3 bucket event for which to invoke the Lambda function. For more information, see [Supported Event Types](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: Yes

*Type*: String

 `Filter`   
The filtering rules that determine which objects invoke the Lambda function. For example, you can create a filter so that only image files with a `.jpg` extension invoke the function when they are added to the S3 bucket.

*Required*: No

*Type*: [Amazon S3 NotificationConfiguration Config Filter](aws-properties-s3-bucket-notificationconfiguration-config-filter.html "Amazon S3 NotificationConfiguration Config Filter")

 `Function`   
The Amazon Resource Name (ARN) of the Lambda function that Amazon S3 invokes when the specified event type occurs.

*Required*: Yes

*Type*: String


