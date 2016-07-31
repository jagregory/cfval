AWS::SNS::TopicPolicy
=====================

The `AWS::SNS::TopicPolicy` resource associates Amazon SNS topics with a policy.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::SNS::TopicPolicy",
  "Properties" :
    {
      "PolicyDocument" : JSON,
      "Topics" : [ List of SNS topic ARNs, ... ]
    }
}
    
```

Properties
----------

 `PolicyDocument`   
A policy document that contains permissions to add to the specified SNS topics.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Topics`   
The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the [Ref function](intrinsic-function-reference-ref.html "Ref") to specify an [AWS::SNS::Topic](aws-properties-sns-topic.html "AWS::SNS::Topic") resource.

*Required*: Yes

*Type*: A list of Amazon SNS topics ARNs

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

For sample `AWS::SNS::TopicPolicy` snippets, see [Declaring an Amazon SNS Topic Policy](quickref-iam.html#scenario-sns-policy "Declaring an Amazon SNS Topic Policy").

