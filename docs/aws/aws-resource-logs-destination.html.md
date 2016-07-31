AWS::Logs::Destination
======================

The `AWS::Logs::Destination` resource creates an Amazon CloudWatch Logs (CloudWatch Logs) destination, which enables you to specify a physical resource (such as an Amazon Kinesis stream) that subscribes to CloudWatch Logs log events from another AWS account. For more information, see [Cross-Account Log Data Sharing with Subscriptions](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CrossAccountSubscriptions.html) in the *Amazon CloudWatch Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Logs::Destination",
  "Properties" : {
    "DestinationName" : String,
    "DestinationPolicy" : String,
    "RoleArn" : String,
    "TargetArn" : String
  }
}
    
```

Properties
----------

 `DestinationName`   
The name of the CloudWatch Logs destination.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DestinationPolicy`   
An AWS Identity and Access Management (IAM) policy that specifies who can write to your destination.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RoleArn`   
The Amazon Resource Name (ARN) of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource (`TargetArn`).

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `TargetArn`   
The ARN of the AWS resource that receives log events. Currently, you can specify only an Amazon Kinesis stream.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name, such as `TestDestination`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

In the following example, the target stream (`TestStream`) can receive log events from the `logger` IAM user that is in the `234567890123` AWS account. The user can call only the `PutSubscriptionFilter` action against the `TestDestination` destination.

``` {.programlisting}
      "DestinationWithName" : {
  "Type" : "AWS::Logs::Destination",
  "Properties" : {
    "DestinationName": "TestDestination",
    "RoleArn": "arn:aws:iam::123456789012:role/LogKinesisRole",
    "TargetArn": "arn:aws:kinesis:us-east-1:123456789012:stream/TestStream",
    "DestinationPolicy": "{\"Version\" : \"2012-10-17\",\"Statement\" : [{\"Effect\" : \"Allow\", \"Principal\" : {\"AWS\" : \"arn:aws:iam::234567890123:user/logger\"},
\"Action\" : \"logs:PutSubscriptionFilter\", \"Resource\" : \"arn:aws:logs:us-east-1:123456789012:destination:TestDestination\"}]}"
  }
}
    
```
