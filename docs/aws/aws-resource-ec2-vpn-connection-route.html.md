AWS::EC2::VPNConnectionRoute
============================

A static route that is associated with a VPN connection between an existing virtual private gateway and a VPN customer gateway. The static route allows traffic to be routed from the virtual private gateway to the VPN customer gateway.

Syntax
------

``` {.programlisting}
      {
   "Type" : "AWS::EC2::VPNConnectionRoute",
   "Properties" : {
      "DestinationCidrBlock" : String
      "VpnConnectionId" : String,
   }
}
    
```

Properties
----------

 `DestinationCidrBlock`   
The CIDR block that is associated with the local subnet of the customer network.

*Required*: Yes.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `VpnConnectionId`   
The ID of the VPN connection.

*Required*: Yes.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

**Example Specifying a static route**

``` {.programlisting}
          "MyConnectionRoute0" : {
   "Type" : "AWS::EC2::VPNConnectionRoute",
   "Properties" : {
      "DestinationCidrBlock" : "10.0.0.0/16",
      "VpnConnectionId" : {"Ref" : "Connection0"}
   }
}
        
```

See Also
--------

-   [CreateVpnConnectionRoute](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnectionRoute.html) in the *Amazon EC2 API Reference*.


