EC2 Network Interface Private IP Specification
==============================================

The `PrivateIpAddressSpecification` type is an embedded property of the [AWS::EC2::NetworkInterface](aws-resource-ec2-network-interface.html "AWS::EC2::NetworkInterface") type.

Syntax
------

``` {.programlisting}
      
{
   "PrivateIpAddress" : String,
   "Primary" : Boolean
}         

    
```

Properties
----------

 `PrivateIpAddress`   
The private IP address of the network interface.

*Required*: Yes

*Type*: String

 `Primary`   
Sets the private IP address as the primary private address. You can set only one primary private IP address. If you don't specify a primary private IP address, Amazon EC2 automatically assigns a primary private IP address.

*Required*: Yes

*Type*: Boolean


