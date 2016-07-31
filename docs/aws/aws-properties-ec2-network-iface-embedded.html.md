EC2 NetworkInterface Embedded Property Type
===========================================

The EC2 Network Interface type is an embedded property of the [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance") type. It specifies a network interface that is to be attached.

Syntax
------

``` {.programlisting}
      
{
   "AssociatePublicIpAddress" : Boolean,
   "DeleteOnTermination" : Boolean,
   "Description" : String,
   "DeviceIndex" : String,
   "GroupSet" : [ String, ... ],
   "NetworkInterfaceId" : String,
   "PrivateIpAddress" : String,
   "PrivateIpAddresses" : [ PrivateIpAddressSpecification, ... ],
   "SecondaryPrivateIpAddressCount" : Integer,
   "SubnetId" : String
}
    
```

Properties
----------

 `AssociatePublicIpAddress`   
Indicates whether the network interface receives a public IP address. You can associate a public IP address with a network interface only if it has a device index of `eth0` and if it is a new network interface (not an existing one). In other words, if you specify true, don't specify a network interface ID. For more information, see [Amazon EC2 Instance IP Addressing](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-instance-addressing.html).

*Required*: No

*Type*: Boolean.

 `DeleteOnTermination`   
Whether to delete the network interface when the instance terminates.

*Required*: No

*Type*: Boolean.

 `Description`   
The description of this network interface.

*Required*: No

*Type*: String

 `DeviceIndex`   
The network interface's position in the attachment order.

*Required*: Yes

*Type*: String

 `GroupSet`   
A list of security group IDs associated with this network interface.

*Required*: No

*Type*: List of strings.

 `NetworkInterfaceId`   
An existing network interface ID.

*Required*: Conditional. If you don't specify the `SubnetId` property, you must specify this property.

*Type*: String

 `PrivateIpAddress`   
Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the `PrivateIpAddresses` property.

*Required*: No

*Type*: String

 `PrivateIpAddresses`   
Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the `Primary` property to `true` in the `PrivateIpAddressSpecification` property. If you want Amazon EC2 to automatically assign private IP addresses, use the `SecondaryPrivateIpCount` property and do not specify this property.

For information about the maximum number of private IP addresses, see [Private IP Addresses Per ENI Per Instance Type](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableIpPerENI) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: list of [PrivateIpAddressSpecification](aws-properties-ec2-network-interface-privateipspec.html "EC2 Network Interface Private IP Specification")

 `SecondaryPrivateIpAddressCount`   
The number of secondary private IP addresses that Amazon EC2 auto assigns to the network interface. Amazon EC2 uses the value of the `PrivateIpAddress` property as the primary private IP address. If you don't specify that property, Amazon EC2 auto assigns both the primary and secondary private IP addresses.

If you want to specify your own list of private IP addresses, use the `PrivateIpAddresses` property and do not specify this property.

For information about the maximum number of private IP addresses, see [Private IP Addresses Per ENI Per Instance Type](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableIpPerENI) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: Integer.

 `SubnetId`   
The ID of the subnet to associate with the network interface.

*Required*: Conditional. If you don't specify the `NetworkInterfaceId` property, you must specify this property.

*Type*: String


