DynamoDB Local Secondary Indexes
================================

Describes local secondary indexes for the [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table") resource. Each index is scoped to a given hash key value. Tables with one or more local secondary indexes are subject to an item collection size limit, where the amount of data within a given item collection cannot exceed 10 GB.

Syntax
------

``` {.programlisting}
      {
  "IndexName" : String,
  "KeySchema" : [ KeySchema, ...],                           
  "Projection" : { Projection }
}
    
```

Properties
----------

 `IndexName`   
The name of the local secondary index. The index name can be 3 – 255 characters long and have no character restrictions.

*Required*: Yes

*Type*: String

 `KeySchema`   
The complete index key schema for the local secondary index, which consists of one or more pairs of attribute names and key types. For local secondary indexes, the hash key must be the same as that of the source table.

*Required*: Yes

*Type*: [DynamoDB Key Schema](aws-properties-dynamodb-keyschema.html "DynamoDB Key Schema")

 `Projection`   
Attributes that are copied (projected) from the source table into the index. These attributes are additions to the primary key attributes and index key attributes, which are automatically projected.

*Required*: Yes

*Type*: [DynamoDB Projection Object](aws-properties-dynamodb-projectionobject.html "DynamoDB Projection Object")

Examples
--------

For an example of a declared local secondary index, see [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table").

