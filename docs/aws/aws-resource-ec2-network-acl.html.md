AWS::EC2::NetworkAcl
====================

Creates a new network ACL in a VPC.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::NetworkAcl",
   "Properties" : {
      "Tags" : [ Resource Tag, ... ],
      "VpcId" : String
   }
}
    
```

Properties
----------

 `Tags`   
An arbitrary set of tags (key–value pairs) for this ACL.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `VpcId`   
The ID of the VPC where the network ACL will be created.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myNetworkAcl" : {
         "Type" : "AWS::EC2::NetworkAcl",
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

-   [CreateNetworkAcl](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAcl.html) in the *Amazon EC2 API Reference*

-   [Network ACLs](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html) in the *Amazon Virtual Private Cloud User Guide*.


