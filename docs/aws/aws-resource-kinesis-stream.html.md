AWS::Kinesis::Stream
====================

Creates an Amazon Kinesis stream that captures and transports data records that are emitted from data sources. For information about creating streams, see [CreateStream](http://docs.aws.amazon.com/kinesis/latest/APIReference/API_CreateStream.html) in the *Amazon Kinesis API Reference*.

Syntax
------

``` {.programlisting}
      {
   "Type" : "AWS::Kinesis::Stream",
   "Properties" : {
      "Name" : String,
      "ShardCount" : Integer,
      "Tags" : [ Resource Tag, ... ]
   }
}
    
```

Properties
----------

 `Name`   
The name of the Amazon Kinesis stream. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the stream name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ShardCount`   
The number of shards that the stream uses. For greater provisioned throughput, increase the number of shards.

*Required*: Yes

*Type*: Integer

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) to associate with the Amazon Kinesis stream.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When you specify an AWS::Kinesis::Stream resource as an argument to the `Ref` function, AWS CloudFormation returns the stream name (physical ID).

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Fn::GetAtt
----------

`Fn::GetAtt` returns a value for the `Arn` attribute.

 `Arn`   
The Amazon resource name (ARN) of the Amazon Kinesis stream, such as `arn:aws:kinesis:us-east-1:123456789012:stream/mystream`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

