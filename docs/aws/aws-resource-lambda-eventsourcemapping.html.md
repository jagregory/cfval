AWS::Lambda::EventSourceMapping
===============================

The `AWS::Lambda::EventSourceMapping` resource specifies a stream as an event source for an AWS Lambda (Lambda) function. The stream can be an Amazon Kinesis stream or an Amazon DynamoDB (DynamoDB) stream. Lambda invokes the associated function when records are posted to the stream. For more information, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the *AWS Lambda Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Lambda::EventSourceMapping",
  "Properties" : {
    "BatchSize" : Integer,
    "Enabled" : Boolean,
    "EventSourceArn" : String,
    "FunctionName" : String,
    "StartingPosition" : String
  }
}
    
```

Properties
----------

 `BatchSize`   
The largest number of records that Lambda retrieves from your event source when invoking your function. Your function receives an event with all the retrieved records. For the default and valid values, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the *AWS Lambda Developer Guide*.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Enabled`   
Indicates whether Lambda begins polling the event source.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EventSourceArn`   
The Amazon Resource Name (ARN) of the Amazon Kinesis or DynamoDB stream that is the source of events. Any record added to this stream can invoke the Lambda function. For more information, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the *AWS Lambda Developer Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `FunctionName`   
The name or ARN of a Lambda function to invoke when Lambda detects an event on the stream.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `StartingPosition`   
The position in the stream where Lambda starts reading. For valid values, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the *AWS Lambda Developer Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following example associates an Amazon Kinesis stream with a Lambda function.

``` {.programlisting}
      "EventSourceMapping": {  
  "Type": "AWS::Lambda::EventSourceMapping",
  "Properties": {
    "EventSourceArn" : { "Fn::Join" : [ "", [ "arn:aws:kinesis:", { "Ref" : "AWS::Region" }, ":", { "Ref" : "AWS::AccountId" }, ":stream/", { "Ref" : "KinesisStream" }] ] },
    "FunctionName" : { "Fn::GetAtt" : ["LambdaFunction", "Arn"] },
    "StartingPosition" : "TRIM_HORIZON"
  }
}
    
```
