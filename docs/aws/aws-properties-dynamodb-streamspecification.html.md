DynamoDB Table StreamSpecification
==================================

`StreamSpecification` is a property of the [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table") resource that defines the settings of a DynamoDB table's stream.

Syntax
------

``` {.programlisting}
      {
   "StreamViewType" : String
}
    
```

Parameters
----------

 `StreamViewType`   
Determines the information that the stream captures when an item in the table is modified. For valid values, see [StreamSpecification](http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_StreamSpecification.html) in the *Amazon DynamoDB API Reference*.

*Required*: Yes

*Type*: String


