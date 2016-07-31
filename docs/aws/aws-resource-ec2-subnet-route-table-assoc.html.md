AWS::EC2::SubnetRouteTableAssociation
=====================================

Associates a subnet with a route table.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::SubnetRouteTableAssociation",
   "Properties" : {
      "RouteTableId" : String,
      "SubnetId" : String,
   }
}     
    
```

Properties
----------

 `RouteTableId`   
The ID of the route table. This is commonly written as a reference to a route table declared elsewhere in the template. For example:

``` {.programlisting}
            "RouteTableId" : { "Ref" : "myRouteTable" }
          
```

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt). However, the physical ID changes when the route table ID is changed.

 `SubnetId`   
The ID of the subnet. This is commonly written as a reference to a subnet declared elsewhere in the template. For example:

``` {.programlisting}
            "SubnetId" : { "Ref" : "mySubnet" }
          
```

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      
{ "Ref": "MyRTA" }
    
```

For the subnet route table association with the logical ID "MyRTA", Ref will return the AWS resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "mySubnetRouteTableAssociation" : {
         "Type" : "AWS::EC2::SubnetRouteTableAssociation",
         "Properties" : {
            "SubnetId" : { "Ref" : "mySubnet" },
            "RouteTableId" : { "Ref" : "myRouteTable" }
         }
      }
   }
}     
    
```

See Also
--------

-   [AssociateRouteTable](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateRouteTable.html) in the *Amazon EC2 API Reference*


