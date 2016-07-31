AWS::EC2::RouteTable
====================

Creates a new route table within a VPC. After you create a new route table, you can add routes and associate the table with a subnet.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::RouteTable",
   "Properties" : {
      "VpcId" : String,
      "Tags" : [ Resource Tag, ... ]
   }
}     
    
```

Properties
----------

 `VpcId`   
The ID of the VPC where the route table will be created.

Example: vpc-11ad4878

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this route table.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Return Values
-------------

### Ref

When you specify an AWS::EC2::RouteTable type as an argument to the `Ref` function, AWS CloudFormation returns the route table ID, such as `rtb-12a34567`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

**Example**

The following example snippet uses the VPC ID from a VPC named *myVPC* that was declared elsewhere in the same template.

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myRouteTable" : {
         "Type" : "AWS::EC2::RouteTable",
         "Properties" : {
            "VpcId" : { "Ref" : "myVPC" },
            "Tags" : [ { "Key" : "foo", "Value" : "bar" } ]
         }
      }
   }
}        
        
```

See Also
--------

-   [AWS::EC2::Route](aws-resource-ec2-route.html "AWS::EC2::Route")

-   [CreateRouteTable](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRouteTable.html) in the *Amazon EC2 API Reference*

-   [Route Tables](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*

-   [Using Tags](http://docs.aws.amazon.com/AWSEC2/latest/DeveloperGuide/Using_Tags.html) in the *Amazon Elastic Compute Cloud User Guide*


