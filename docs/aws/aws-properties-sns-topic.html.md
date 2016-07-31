AWS::SNS::Topic
===============

The `AWS::SNS::Topic` type creates an Amazon Simple Notification Service (Amazon SNS) topic.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::SNS::Topic",
  "Properties" : {
    "DisplayName" : String,
    "Subscription" : [ SNS Subscription, ... ],
    "TopicName" : String
  }
}
    
```

Properties
----------

 `DisplayName`   
A developer-defined string that can be used to identify this SNS topic.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Subscription`   
The SNS subscriptions (endpoints) for this topic.

*Required*: No

*Type*: List of [SNS Subscriptions](aws-properties-sns-subscription.html "Amazon SNS Subscription Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `TopicName`   
A name for the topic. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

For the `AWS::SNS::Topic` resource, the `Ref` intrinsic function returns the topic ARN, for example: `arn:aws:sns:us-east-1:123456789012:mystack-mytopic-NZJ5JSMVGFIE`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `TopicName`   
Returns the name for an Amazon SNS topic.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

An example of an SNS topic subscribed to by two SQS queues:

``` {.programlisting}
      
"MySNSTopic" : {
   "Type" : "AWS::SNS::Topic",
   "Properties" : {
      "Subscription" : [
         { "Endpoint" : { "Fn::GetAtt" : [ "MyQueue1", "Arn" ] }, "Protocol" : "sqs" },
         { "Endpoint" : { "Fn::GetAtt" : [ "MyQueue2", "Arn" ] }, "Protocol" : "sqs" }
      ],
      "TopicName" : "SampleTopic"
   }
}
    
```

See Also
--------

-   [Using an AWS CloudFormation Template to Create a Topic that Sends Messages to Amazon SQS Queues](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToSQS.cloudformation.html) in the *Amazon Simple Notification Service Developer Guide*


