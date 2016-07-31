DynamoDB Provisioned Throughput
===============================

Describes a set of provisioned throughput values for an [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table") resource. DynamoDB uses these capacity units to allocate sufficient resources to provide the requested throughput.

For a complete discussion of DynamoDB provisioned throughput values, see [Specifying Read and Write Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the *DynamoDB Developer Guide*.

Syntax
------

``` {.programlisting}
      {
   "ReadCapacityUnits" : Number,
   "WriteCapacityUnits" : Number
}
    
```

Parameters
----------

 `ReadCapacityUnits`   
Sets the desired minimum number of consistent reads of items (up to 1KB in size) per second for the specified table before Amazon DynamoDB balances the load.

*Required*: Yes

*Type*: Number

 `WriteCapacityUnits`   
Sets the desired minimum number of consistent writes of items (up to 1KB in size) per second for the specified table before Amazon DynamoDB balances the load.

*Required*: Yes

*Type*: Number

Note

For detailed information about the limits of provisioned throughput values in DynamoDB, see [Limits in Amazon DynamoDB](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *DynamoDB Developer Guide*.

