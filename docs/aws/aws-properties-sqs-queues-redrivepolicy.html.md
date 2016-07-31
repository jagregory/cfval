Amazon SQS RedrivePolicy
========================

The RedrivePolicy type is a property of the [AWS::SQS::Queue](aws-properties-sqs-queues.html "AWS::SQS::Queue") resource.

Syntax
------

``` {.programlisting}
      {
  "deadLetterTargetArn" : String,
  "maxReceiveCount" : Integer
}
    
```

Properties
----------

 `deadLetterTargetArn`   
The Amazon Resource Name (ARN) of the dead letter queue to which the messages are sent to after the `maxReceiveCount` value has been exceeded.

*Required*: No

*Type*: String

 `maxReceiveCount`   
The number of times a message is delivered to the source queue before being sent to the dead letter queue.

*Required*: No

*Type*: Integer


