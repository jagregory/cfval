DynamoDB Global Secondary Indexes
=================================

Describes global secondary indexes for the [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table") resource.

Syntax
------

``` {.programlisting}
      {
  "IndexName" : String,
  "KeySchema" : [ KeySchema, ... ],
  "Projection" : { Projection },
  "ProvisionedThroughput" : { ProvisionedThroughput }
}
    
```

Properties
----------

 `IndexName`   
The name of the global secondary index. The index name can be 3 – 255 characters long and have no character restrictions.

*Required*: Yes

*Type*: String

 `KeySchema`   
The complete index key schema for the global secondary index, which consists of one or more pairs of attribute names and key types.

*Required*: Yes

*Type*: [DynamoDB Key Schema](aws-properties-dynamodb-keyschema.html "DynamoDB Key Schema")

 `Projection`   
Attributes that are copied (projected) from the source table into the index. These attributes are in addition to the primary key attributes and index key attributes, which are automatically projected.

*Required*: Yes

*Type*: [DynamoDB Projection Object](aws-properties-dynamodb-projectionobject.html "DynamoDB Projection Object")

 `ProvisionedThroughput`   
The provisioned throughput settings for the index.

*Required*: Yes

*Type*: [DynamoDB Provisioned Throughput](aws-properties-dynamodb-provisionedthroughput.html "DynamoDB Provisioned Throughput")


