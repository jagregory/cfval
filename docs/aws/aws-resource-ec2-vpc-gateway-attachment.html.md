AWS::EC2::VPCGatewayAttachment
==============================

Attaches a gateway to a VPC.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::VPCGatewayAttachment",
   "Properties" : {
      "InternetGatewayId" : String,
      "VpcId" : String,
      "VpnGatewayId" : String
   }
}     
    
```

Properties
----------

 `InternetGatewayId`   
The ID of the Internet gateway.

*Required*: Conditional You must specify either InternetGatewayId or VpnGatewayId, but not both.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VpcId`   
The ID of the VPC to associate with this gateway.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VpnGatewayId`   
The ID of the virtual private network (VPN) gateway to attach to the VPC.

*Required*: Conditional You must specify either InternetGatewayId or VpnGatewayId, but not both.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

**Example Attaching both an Internet gateway and a VPN gateway to a VPC**

To attach both an Internet gateway and a VPN gateway to a VPC, you must specify two separate AWS::EC2::VPCGatewayAttachment resources:

``` {.programlisting}
          
"AttachGateway" : {
   "Type" : "AWS::EC2::VPCGatewayAttachment",
   "Properties" : {
      "VpcId" : { "Ref" : "VPC" },
      "InternetGatewayId" : { "Ref" : "myInternetGateway" }
   }
},

"AttachVpnGateway" : {
   "Type" : "AWS::EC2::VPCGatewayAttachment",
   "Properties" : {
      "VpcId" : { "Ref" : "VPC" },
      "VpnGatewayId" : { "Ref" : "myVPNGateway" }
   }
},       
        
```

See Also
--------

-   [AttachVpnGateway](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVpnGateway.html) in the *Amazon EC2 API Reference*.


