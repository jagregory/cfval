AWS::SQS::QueuePolicy
=====================

The AWS::SQS::QueuePolicy type applies a policy to SQS queues.

AWS::SQS::QueuePolicy Snippet: [Declaring an Amazon SQS Policy](quickref-iam.html#scenario-sqs-policy "Declaring an Amazon SQS Policy")

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::SQS::QueuePolicy",
   "Properties" : {
      "PolicyDocument" : JSON,
      "Queues" : [ String, ... ]
   }
}
    
```

Properties
----------

 `PolicyDocument`   
A policy document containing permissions to add to the specified SQS queues.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Queues`   
The URLs of the queues to which you want to add the policy. You can use the [Ref function](intrinsic-function-reference-ref.html "Ref") to specify an [AWS::SQS::Queue](aws-properties-sqs-queues.html "AWS::SQS::Queue") resource.

*Required*: Yes

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)


