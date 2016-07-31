AWS::EC2::VPNConnection
=======================

Creates a new VPN connection between an existing virtual private gateway and a VPN customer gateway.

For more information, go to [CreateVpnConnection](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnection.html) in the *Amazon EC2 API Reference*.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::VPNConnection",
   "Properties" : {
      "Type" : String,
      "CustomerGatewayId" : GatewayID,
      "StaticRoutesOnly" : Boolean,
      "Tags" :  [ Resource Tag, ... ],
      "VpnGatewayId" : GatewayID
   }
} 
    
```

Properties
----------

 `Type`   
The type of VPN connection this virtual private gateway supports.

Example: "ipsec.1"

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `CustomerGatewayId`   
The ID of the customer gateway. This can either be an embedded JSON object or a reference to a Gateway ID.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `StaticRoutesOnly`   
Indicates whether the VPN connection requires static routes.

*Required*: Conditional: If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify `true`.

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
The tags that you want to attach to the resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `VpnGatewayId`   
The ID of the virtual private gateway. This can either be an embedded JSON object or a reference to a Gateway ID.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "MyVPNConnection" }
    
```

For the VPNConnection with the logical ID "MyVPNConnection", `Ref` will return the VPN connection's resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Examples
-----------------

**Example VPNConnection**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myVPNConnection" : {
         "Type" : "AWS::EC2::VPNConnection",
         "Properties" : {
            "Type" : "ipsec.1",
            "StaticRoutesOnly" : "true",
            "CustomerGatewayId" : {"Ref" : "myCustomerGateway"},
            "VpnGatewayId" : {"Ref" : "myVPNGateway"}
         }
      }
   }
}
        
```


