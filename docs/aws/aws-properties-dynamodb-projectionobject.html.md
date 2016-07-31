DynamoDB Projection Object
==========================

Attributes that are copied (projected) from the source table into the index. These attributes are additions to the primary key attributes and index key attributes, which are automatically projected.

Syntax
------

``` {.programlisting}
      {
  "NonKeyAttributes" : [ String, ... ],
  "ProjectionType" : String
}
    
```

Properties
----------

 `NonKeyAttributes`   
The non-key attribute names that are projected into the index.

For local secondary indexes, the total count of `NonKeyAttributes` summed across all of the local secondary indexes must not exceed 20. If you project the same attribute into two different indexes, this counts as two distinct attributes in determining the total.

*Required*: No

*Type*: List of strings

 `ProjectionType`   
The set of attributes that are projected into the index:

 `KEYS_ONLY`   
Only the index and primary keys are projected into the index.

 `INCLUDE`   
Only the specified table attributes are projected into the index. The list of projected attributes are in `NonKeyAttributes`.

 `ALL`   
All of the table attributes are projected into the index.

*Required*: No

*Type*: String


