AWS::RDS::DBSubnetGroup
=======================

The AWS::RDS::DBSubnetGroup type creates an RDS database subnet group. Subnet groups must contain at least two subnet in two different Availability Zones in the same region.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::RDS::DBSubnetGroup",
   "Properties" : {
      "DBSubnetGroupDescription" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Resource Tag, ... ]
   }
}     
    
```

Properties
----------

 `DBSubnetGroupDescription`   
The description for the DB Subnet Group.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SubnetIds`   
The EC2 Subnet IDs for the DB Subnet Group.

*Required*: Yes

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Tags`   
The tags that you want to attach to the RDS database subnet group.

*Required*: No

*Type*: A list of [resource tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

### Ref

When you pass the logical ID of an `AWS::RDS::DBSubnetGroup` resource to the intrinsic `Ref` function, the function returns the name of the DB subnet group, such as `mystack-mydbsubnetgroup-0a12bc456789de0fg`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myDBSubnetGroup" : {
         "Type" : "AWS::RDS::DBSubnetGroup",
         "Properties" : {
            "DBSubnetGroupDescription" : "description",
            "SubnetIds" : [ "subnet-7b5b4112", "subnet-7b5b4115" ],
            "Tags" : [ {"key" : "value", "key2" : "value2"} ]
         }
      }
   }
}     
    
```

See Also
--------

-   [CreateDBSubnetGroup](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBSubnetGroup.html) in the *Amazon Relational Database Service API Reference*

-   [ModifyDBSubnetGroup](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBSubnetGroup.html) in the *Amazon Relational Database Service API Reference*

-   [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates")


