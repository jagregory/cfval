AWS::EC2::VPNGatewayRoutePropagation
====================================

Enables a virtual private gateway (VGW) to propagate routes to the routing tables of a VPC.

Note

If you reference a VPN gateway that is in the same template as your VPN gateway route propagation, you must explicitly declare a dependency on the VPN gateway attachment. The `AWS::EC2::VPNGatewayRoutePropagation` resource cannot use the VPN gateway until it has successfully attached to the VPC. Add a [DependsOn](aws-attribute-dependson.html "DependsOn Attribute") attribute in the `AWS::EC2::VPNGatewayRoutePropagation` resource to explicitly declare a dependency on the VPN gateway attachment.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::VPNGatewayRoutePropagation",
   "Properties" : {
      "RouteTableIds" : [ String, ... ],
      "VpnGatewayId" : String
   }
}
     
    
```

Properties
----------

 `RouteTableIds`   
A list of routing table IDs that are associated with a VPC. The routing tables must be associated with the same VPC that the virtual private gateway is attached to.

*Required*: Yes

*Type*: List of route table IDs

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VpnGatewayId`   
The ID of the virtual private gateway that is attached to a VPC. The virtual private gateway must be attached to the same VPC that the routing tables are associated with.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "myVPNGatewayRouteProp" }
    
```

For the VPN gateway with the logical ID `myVPNGatewayRouteProp`, `Ref` will return the gateway's resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      "myVPNGatewayRouteProp" : {
  "Type" : "AWS::EC2::VPNGatewayRoutePropagation",
  "Properties" : {
    "RouteTableIds" : [{"Ref" : "PrivateRouteTable"}],
    "VpnGatewayId" : {"Ref" : "VPNGateway"}
  }
}
    
```

See Also
--------

-   [EnableVgwRoutePropagation](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVgwRoutePropagation.html) in the *Amazon EC2 API Reference*.


