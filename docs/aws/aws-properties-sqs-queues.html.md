AWS::SQS::Queue
===============

The AWS::SQS::Queue type creates an Amazon SQS queue.

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::SQS::Queue",
   "Properties": {
      "DelaySeconds": Integer,
      "MaximumMessageSize": Integer,
      "MessageRetentionPeriod": Integer,
      "QueueName": String,
      "ReceiveMessageWaitTimeSeconds": Integer,
      "RedrivePolicy": RedrivePolicy,
      "VisibilityTimeout": Integer
   }
}     
    
```

Properties
----------

 `DelaySeconds`   
The time in seconds that the delivery of all messages in the queue will be delayed. You can specify an integer value of `0` to `900` (15 minutes). The default value is `0`.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MaximumMessageSize`   
The limit of how many bytes a message can contain before Amazon SQS rejects it. You can specify an integer value from `1024` bytes (1 KiB) to `262144` bytes (256 KiB). The default value is `262144` (256 KiB).

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MessageRetentionPeriod`   
The number of seconds Amazon SQS retains a message. You can specify an integer value from `60` seconds (1 minute) to `1209600` seconds (14 days). The default value is `345600` seconds (4 days).

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `QueueName`   
A name for the queue. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the queue name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ReceiveMessageWaitTimeSeconds`   
Specifies the duration, in seconds, that the `ReceiveMessage` action call waits until a message is in the queue in order to include it in the response, as opposed to returning an empty response if a message is not yet available. You can specify an integer from `1` to `20`. The short polling is used as the default or when you specify `0` for this property. For more information, see [Amazon SQS Long Poll](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html).

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RedrivePolicy`   
Specifies an existing dead letter queue to receive messages after the source queue (this queue) fails to process a message a specified number of times.

*Required*: No

*Type*: [Amazon SQS RedrivePolicy](aws-properties-sqs-queues-redrivepolicy.html "Amazon SQS RedrivePolicy")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VisibilityTimeout`   
The length of time during which the queue will be unavailable once a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.

Values must be from 0 to 43200 seconds (12 hours). If no value is specified, the default value of 30 seconds will be used.

For more information about SQS Queue visibility timeouts, see [Visibility Timeout](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html) in the *Amazon Simple Queue Service Developer Guide*.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

The AWS::SQS::Queue type returns the queue URL, for example: `https://sqs.us-east-1.amazonaws.com/123456789012/aa4-MyQueue-Z5NOSZO2PZE9`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
Returns the Amazon Resource Name (ARN) of the queue. For example: `arn:aws:sqs:us-east-1:123456789012:mystack-myqueue-15PG5C2FC1CW8`

 `QueueName`   
Returns the queue name. For example:

mystack-myqueue-1VF9BKQH5BJVI

Examples
--------

### SQS Queue with Cloudwatch Alarms

``` {.programlisting}
        {
  "AWSTemplateFormatVersion" : "2010-09-09",
 
  "Description" : "AWS CloudFormation Sample Template SQS_With_CloudWatch_Alarms: Sample template showing how to create an SQS queue with Amazon CloudWatch alarms on queue depth. **WARNING** This template creates an Amazon SQS queue and one or more Amazon CloudWatch alarms. You will be billed for the AWS resources used if you create a stack from this template.",
 
  "Parameters" : {
    "AlarmEmail": {
      "Default": "nobody@amazon.com",
      "Description": "Email address to notify if operational problems arise",
      "Type": "String"
    }
  },
 
  "Resources" : {
    "MyQueue" : {
      "Type" : "AWS::SQS::Queue",
      "Properties" : {
         "QueueName" : "SampleQueue"
      }
    },
    "AlarmTopic": {
      "Type": "AWS::SNS::Topic",
      "Properties": {
        "Subscription": [{
          "Endpoint": { "Ref": "AlarmEmail" },
          "Protocol": "email"
        }]
      }
    },
    "QueueDepthAlarm": {
      "Type": "AWS::CloudWatch::Alarm",
      "Properties": {
        "AlarmDescription": "Alarm if queue depth grows beyond 10 messages",
        "Namespace": "AWS/SQS",
        "MetricName": "ApproximateNumberOfMessagesVisible",
        "Dimensions": [{
          "Name": "QueueName",
          "Value" : { "Fn::GetAtt" : ["MyQueue", "QueueName"] }
        }],
        "Statistic": "Sum",
        "Period": "300",
        "EvaluationPeriods": "1",
        "Threshold": "10",
        "ComparisonOperator": "GreaterThanThreshold",
        "AlarmActions": [{
          "Ref": "AlarmTopic"
        }],
        "InsufficientDataActions": [{
          "Ref": "AlarmTopic"
        }]
      }
    }
  },
  "Outputs" : {
    "QueueURL" : {
      "Description" : "URL of newly created SQS Queue",
      "Value" : { "Ref" : "MyQueue" }
    },
    "QueueARN" : {
      "Description" : "ARN of newly created SQS Queue",
      "Value" : { "Fn::GetAtt" : ["MyQueue", "Arn"]}
    },
    "QueueName" : {
      "Description" : "Name newly created SQS Queue",
      "Value" : { "Fn::GetAtt" : ["MyQueue", "QueueName"]}
    }
  }
}
      
```

### SQS Queue with a Dead Letter Queue

The following sample creates a source queue and a dead letter queue. Because the source queue specifies the dead letter queue in its redrive policy, the source queue is dependent on the creation of the dead letter queue.

``` {.programlisting}
        {
  "AWSTemplateFormatVersion" : "2010-09-09",
   
  "Resources" : {
    "MySourceQueue" : {
      "Type" : "AWS::SQS::Queue",
      "Properties" : {
        "RedrivePolicy": {
          "deadLetterTargetArn" : {"Fn::GetAtt" : [ "MyDeadLetterQueue" , "Arn" ]},
          "maxReceiveCount" : 5
        }
      }
    },    
    "MyDeadLetterQueue" : {
      "Type" : "AWS::SQS::Queue"
    }
  },

  "Outputs" : {
    "SourceQueueURL" : {
      "Description" : "URL of the source queue",
      "Value" : { "Ref" : "MySourceQueue" }
    },
    "SourceQueueARN" : {
      "Description" : "ARN of the source queue",
      "Value" : { "Fn::GetAtt" : ["MySourceQueue", "Arn"]}
    },
    "DeadLetterQueueURL" : {
      "Description" : "URL of the dead letter queue",
      "Value" : { "Ref" : "MyDeadLetterQueue" }
    },
    "DeadLetterQueueARN" : {
      "Description" : "ARN of the dead letter queue",
      "Value" : { "Fn::GetAtt" : ["MyDeadLetterQueue", "Arn"]}
    }    
  }
}
      
```

See Also
--------

-   [CreateQueue](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/Query_QueryCreateQueue.html) in the Amazon Simple Queue Service API Reference

-   [What is Amazon Simple Queue Service?](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/Welcome.html) in the Amazon Simple Queue Service Developer Guide


