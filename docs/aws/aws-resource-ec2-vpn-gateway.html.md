AWS::EC2::VPNGateway
====================

Creates a virtual private gateway. A virtual private gateway is the VPC-side endpoint for your VPN connection.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::VPNGateway",
   "Properties" : {
      "Type" : String,
      "Tags" : [ Resource Tag, ... ]
   }
}
     
    
```

Properties
----------

 `Type`   
The type of VPN connection this virtual private gateway supports. The only valid value is `"ipsec.1"`.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "MyVPNGateway" }
    
```

For the VPN gateway with the logical ID "MyVPNGateway", `Ref` will return the gateway's resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myVPNGateway" : {
         "Type" : "AWS::EC2::VPNGateway",
         "Properties" : {
            "Type" : "ipsec.1",
            "Tags" : [ { "Key" : "Use", "Value" : "Test" } ]
         }
      }
   }
}        
    
```

See Also
--------

-   [CreateVpnGateway](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnGateway.html) in the *Amazon EC2 API Reference*.


