AWS::DynamoDB::Table
====================

Creates a DynamoDB table.

Note

AWS CloudFormation typically creates DynamoDB tables in parallel. However, if your template includes multiple DynamoDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DynamoDB limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DynamoDB returns an error and the stack operation fails. For a sample snippet, see [DynamoDB Table with a DependsOn Attribute](aws-resource-dynamodb-table.html#cfn-dynamodb-table-examples-dependson "DynamoDB Table with a DependsOn Attribute").

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::DynamoDB::Table",
  "Properties" : {
    "AttributeDefinitions" : [ AttributeDefinitions, ... ],
    "GlobalSecondaryIndexes" : [ GlobalSecondaryIndexes, ... ],
    "KeySchema" : [ KeySchema, ... ],
    "LocalSecondaryIndexes" : [ LocalSecondaryIndexes, ... ],
    "ProvisionedThroughput" : ProvisionedThroughput,
    "StreamSpecification" : StreamSpecification,
    "TableName" : String
  }
}
    
```

Properties
----------

 `AttributeDefinitions`   
A list of `AttributeName` and `AttributeType` objects that describe the key schema for the table and indexes.

*Required*: Yes

*Type*: List of [DynamoDB Attribute Definitions](aws-properties-dynamodb-attributedef.html "DynamoDB Attribute Definitions")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `GlobalSecondaryIndexes`   
Global secondary indexes to be created on the table. You can create up to 5 global secondary indexes.

Important

If you update a table to include a new global secondary index, AWS CloudFormation initiates the index creation and then proceeds with the stack update. AWS CloudFormation doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You cannot use the index or update the table until the index's status is `ACTIVE`. You can track its status by using the DynamoDB [`DescribeTable`](http://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.

If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index.

*Required*: No

*Type*: List of [DynamoDB Global Secondary Indexes](aws-properties-dynamodb-gsi.html "DynamoDB Global Secondary Indexes")

*Update requires*: Updates are not supported. with the following exceptions:

-   If you update only the provisioned throughput values of global secondary indexes, you can update the table [without interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

-   You can delete or add one global secondary index [without interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt). If you do both in the same update (for example, by changing the index's logical ID), the update fails.

 `KeySchema`   
Specifies the attributes that make up the primary key for the table. The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.

*Required*: Yes

*Type*: List of [DynamoDB Key Schema](aws-properties-dynamodb-keyschema.html "DynamoDB Key Schema")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `LocalSecondaryIndexes`   
Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.

*Required*: No

*Type*: List of [DynamoDB Local Secondary Indexes](aws-properties-dynamodb-lsi.html "DynamoDB Local Secondary Indexes")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ProvisionedThroughput`   
Throughput for the specified table, consisting of values for ReadCapacityUnits and WriteCapacityUnits. For more information about the contents of a provisioned throughput structure, see [DynamoDB Provisioned Throughput](aws-properties-dynamodb-provisionedthroughput.html "DynamoDB Provisioned Throughput").

*Required*: Yes

*Type*: [DynamoDB Provisioned Throughput](aws-properties-dynamodb-provisionedthroughput.html "DynamoDB Provisioned Throughput")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `StreamSpecification`   
The settings for the DynamoDB table stream, which capture changes to items stored in the table.

*Required*: No

*Type*: [DynamoDB Table StreamSpecification](aws-properties-dynamodb-streamspecification.html "DynamoDB Table StreamSpecification")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) to the table; however, the stream is replaced.

 `TableName`   
A name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide*.

Return Value
------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
        { "Ref": "MyResource" }
      
```

For the resource with the logical ID `myDynamoDBTable`, `Ref` will return the DynamoDB table name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `StreamArn`   
The Amazon Resource Name (ARN) of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

DynamoDB Table with Local and Secondary Indexes
-----------------------------------------------

The following sample creates an DynamoDB table with `Album`, `Artist`, `Sales`, `NumberOfSongs` as attributes. The primary key includes the `Album` attribute as the hash key and `Artist` attribute as the range key. The table also includes two global and one secondary index. For querying the number of sales for a given artist, the global secondary index uses the `Sales` attribute as the hash key and the `Artist` attribute as the range key.

For querying the sales based on the number of songs, the global secondary index uses the `NumberOfSongs` attribute as the hash key and the `Sales` attribute as the range key.

For querying the sales of an album, the local secondary index uses the same hash key as the table but uses the `Sales` attribute as the range key.

``` {.programlisting}
      {
  "AWSTemplateFormatVersion" : "2010-09-09",
  "Resources" : {
    "myDynamoDBTable" : {
      "Type" : "AWS::DynamoDB::Table",
      "Properties" : {
        "AttributeDefinitions" : [
          {
            "AttributeName" : "Album",
            "AttributeType" : "S"   
          },
          {
            "AttributeName" : "Artist",
            "AttributeType" : "S"
          },
          {
            "AttributeName" : "Sales",
            "AttributeType" : "N"
          },
          {
            "AttributeName" : "NumberOfSongs",
            "AttributeType" : "N"
          }
        ],
        "KeySchema" : [
          {
            "AttributeName" : "Album",
            "KeyType" : "HASH"
          },
          {
            "AttributeName" : "Artist",
            "KeyType" : "RANGE"
          }
        ],
        "ProvisionedThroughput" : {
          "ReadCapacityUnits" : "5",
          "WriteCapacityUnits" : "5"
        },
        "TableName" : "myTableName",
        "GlobalSecondaryIndexes" : [{
          "IndexName" : "myGSI",
          "KeySchema" : [
            {
              "AttributeName" : "Sales",
              "KeyType" : "HASH"
            },
            {
              "AttributeName" : "Artist",
              "KeyType" : "RANGE"
            }
          ],                         
          "Projection" : {
            "NonKeyAttributes" : ["Album","NumberOfSongs"],
            "ProjectionType" : "INCLUDE"
          },
          "ProvisionedThroughput" : {
            "ReadCapacityUnits" : "5",
            "WriteCapacityUnits" : "5"
          }
        },
        {
          "IndexName" : "myGSI2",
          "KeySchema" : [
            {
              "AttributeName" : "NumberOfSongs",
              "KeyType" : "HASH"
            },
            {
              "AttributeName" : "Sales",
              "KeyType" : "RANGE"
            }
          ],                         
          "Projection" : {
            "NonKeyAttributes" : ["Album","Artist"],
            "ProjectionType" : "INCLUDE"
          },
          "ProvisionedThroughput" : {
            "ReadCapacityUnits" : "5",
            "WriteCapacityUnits" : "5"
          }
        }],
        "LocalSecondaryIndexes" :[{
          "IndexName" : "myLSI",
          "KeySchema" : [
            {
              "AttributeName" : "Album",
              "KeyType" : "HASH"
            },
            {
              "AttributeName" : "Sales",
              "KeyType" : "RANGE"
            }
          ],                           
          "Projection" : {
            "NonKeyAttributes" : ["Artist","NumberOfSongs"],
            "ProjectionType" : "INCLUDE"
          }
        }]
      }
    }
  }
}
    
```

DynamoDB Table with a DependsOn Attribute
-----------------------------------------

If you include multiple DynamoDB tables with indexes in a single template, you must include dependencies so that the tables are created sequentially. DynamoDB limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DynamoDB returns an error and the stack operation fails.

The following sample assumes that the `myFirstDDBTable` table is declared in the same template as the `mySecondDDBTable` table, and both tables include a secondary index. The `mySecondDDBTable` table includes a dependency on the `myFirstDDBTable` table so that AWS CloudFormation creates the tables one at a time.

``` {.programlisting}
      "mySecondDDBTable" : {
  "Type" : "AWS::DynamoDB::Table",
  "DependsOn" : "myFirstDDBTable" ,
  "Properties" : {
    "AttributeDefinitions" : [
      {
        "AttributeName" : "ArtistId",
        "AttributeType" : "S"        
      },
      {
        "AttributeName" : "Concert",
        "AttributeType" : "S"        
      },
      {
        "AttributeName" : "TicketSales",
        "AttributeType" : "S"        
      }
    ],
    "KeySchema" : [
      {
        "AttributeName" : "ArtistId",
        "KeyType" : "HASH"
      },
      {
        "AttributeName" : "Concert",
        "KeyType" : "RANGE"
      }
    ],
    "ProvisionedThroughput" : {
      "ReadCapacityUnits" : {"Ref" : "ReadCapacityUnits"},
      "WriteCapacityUnits" : {"Ref" : "WriteCapacityUnits"}
    },
    "GlobalSecondaryIndexes" : [{
      "IndexName" : "myGSI",
      "KeySchema" : [
        {
          "AttributeName" : "TicketSales",
          "KeyType" : "HASH"
        }
      ],                           
      "Projection" : {
        "ProjectionType" : "KEYS_ONLY"
      },    
      "ProvisionedThroughput" : {
        "ReadCapacityUnits" : {"Ref" : "ReadCapacityUnits"},
        "WriteCapacityUnits" : {"Ref" : "WriteCapacityUnits"}
      }
    }]
  }
}
    
```
