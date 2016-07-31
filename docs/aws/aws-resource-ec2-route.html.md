AWS::EC2::Route
===============

Creates a new route in a route table within a VPC. The route's target can be either a gateway attached to the VPC or a NAT instance in the VPC.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EC2::Route",
  "Properties" : {
    "DestinationCidrBlock" : String,
    "GatewayId" : String,
    "InstanceId" : String,
    "NatGatewayId" : String,
    "NetworkInterfaceId" : String,
    "RouteTableId" : String,
    "VpcPeeringConnectionId" : String
  }
}
    
```

Properties
----------

 `DestinationCidrBlock`   
The CIDR address block used for the destination match. For example, `0.0.0.0/0`. Routing decisions are based on the most specific match.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `GatewayId`   
The ID of an Internet gateway or virtual private gateway that is attached to your VPC. For example: `igw-eaad4883`.

For route entries that specify a gateway, you must specify a dependency on the gateway attachment resource. For more information, see [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

*Required*: Conditional. You must specify only one of the following properties: `GatewayId`, `InstanceId`, `NatGatewayId`, `NetworkInterfaceId`, or `VpcPeeringConnectionId`.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `InstanceId`   
The ID of a NAT instance in your VPC. For example, `i-1a2b3c4d`.

*Required*: Conditional. You must specify only one of the following properties: `GatewayId`, `InstanceId`, `NatGatewayId`, `NetworkInterfaceId`, or `VpcPeeringConnectionId`.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NatGatewayId`   
The ID of a NAT gateway. For example, `nat-0a12bc456789de0fg`.

*Required*: Conditional. You must specify only one of the following properties: `GatewayId`, `InstanceId`, `NatGatewayId`, `NetworkInterfaceId`, or `VpcPeeringConnectionId`.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NetworkInterfaceId`   
Allows the routing of network interface IDs.

*Required*: Conditional. You must specify only one of the following properties: `GatewayId`, `InstanceId`, `NatGatewayId`, `NetworkInterfaceId`, or `VpcPeeringConnectionId`.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RouteTableId`   
The ID of the [route table](aws-resource-ec2-route-table.html "AWS::EC2::RouteTable") where the route will be added.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `VpcPeeringConnectionId`   
The ID of a VPC peering connection.

*Required*: Conditional. You must specify only one of the following properties: `GatewayId`, `InstanceId`, `NatGatewayId`, `NetworkInterfaceId`, or `VpcPeeringConnectionId`.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

**Example Route with Gateway ID**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myRoute" : {
         "Type" : "AWS::EC2::Route",
         "DependsOn" : "GatewayToInternet",
         "Properties" : {
            "RouteTableId" : { "Ref" : "myRouteTable" },
            "DestinationCidrBlock" : "0.0.0.0/0",
            "GatewayId" : { "Ref" : "myInternetGateway" }
         }
      }
   }
}        
        
```

**Example Route with Instance ID**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myRoute" : {
         "Type" : "AWS::EC2::Route",
         "Properties" : {
            "RouteTableId" : { "Ref" : "myRouteTable" },
            "DestinationCidrBlock" : "0.0.0.0/0",
            "InstanceId" : { "Ref" : "myInstance" }
         }
      }
   }
}        
        
```

**Example Route with Network Interface ID.**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myRoute" : {
         "Type" : "AWS::EC2::Route",
         "Properties" : {
            "RouteTableId" : { "Ref" : "myRouteTable" },
            "DestinationCidrBlock" : "0.0.0.0/0",
            "NetworkInterfaceId" : { "Ref" : "eni-1a2b3c4d" }
         }
      }
   }
}        
        
```

**Example Route with VPC peering connection ID.**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myRoute" : {
         "Type" : "AWS::EC2::Route",
         "Properties" : {
            "RouteTableId" : { "Ref" : "myRouteTable" },
            "DestinationCidrBlock" : "0.0.0.0/0",
            "VpcPeeringConnectionId" : { "Ref" : "myVPCPeeringConnectionID" }
         }
      }
   }
}        
        
```

See Also
--------

-   [AWS::EC2::RouteTable](aws-resource-ec2-route-table.html "AWS::EC2::RouteTable")

-   [CreateRoute](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRoute.html) in the *Amazon EC2 API Reference*

-   [Route Tables](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*.


