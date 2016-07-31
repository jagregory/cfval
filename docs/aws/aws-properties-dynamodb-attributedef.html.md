DynamoDB Attribute Definitions
==============================

A list of attribute definitions for the [AWS::DynamoDB::Table](aws-resource-dynamodb-table.html "AWS::DynamoDB::Table") resource. Each element is composed of an `AttributeName` and `AttributeType`.

Syntax
------

``` {.programlisting}
      {
  "AttributeName" : String,
  "AttributeType" : String
}
    
```

Properties
----------

 `AttributeName`   
The name of an attribute. Attribute names can be 1 – 255 characters long and have no character restrictions.

*Required*: Yes

*Type*: String

 `AttributeType`   
The data type for the attribute. You can specify `S` for string data, `N` for numeric data, or `B` for binary data.

*Required*: Yes

*Type*: String


