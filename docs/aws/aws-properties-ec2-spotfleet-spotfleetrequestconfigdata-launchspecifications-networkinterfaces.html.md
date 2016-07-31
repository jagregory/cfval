Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces
========================================================================================================

`NetworkInterfaces` is a property of the [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications") property that defines the network interface of the instances.

Syntax
------

``` {.programlisting}
      {
  "AssociatePublicIpAddress" : Boolean,
  "DeleteOnTermination" : Boolean,
  "Description" : String,
  "DeviceIndex" : Integer,
  "Groups" : [ String, ... ],
  "NetworkInterfaceId" : String,
  "PrivateIpAddresses" : [ PrivateIpAddresses, ... ],
  "SecondaryPrivateIpAddressCount" : Integer,
  "SubnetId" : String
}
    
```

Properties
----------

 `AssociatePublicIpAddress`   
Indicates whether monitoring is enabled for the instances.

*Required*: No

*Type*: Boolean

 `DeleteOnTermination`   
Indicates whether to delete the network interface when the instance terminates.

*Required*: No

*Type*: Boolean

 `Description`   
The description of this network interface.

*Required*: No

*Type*: String

 `DeviceIndex`   
The network interface's position in the attachment order.

*Required*: Yes

*Type*: Integer

 `Groups`   
A list of security group IDs to associate with this network interface.

*Required*: No

*Type*: List of strings

 `NetworkInterfaceId`   
A network interface ID.

*Required*: No

*Type*: String

 `PrivateIpAddresses`   
One or more private IP addresses to assign to the network interface. You can designate only one private IP address as primary.

*Required*: No

*Type*: List of [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces PrivateIpAddresses](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces PrivateIpAddresses")

 `SecondaryPrivateIpAddressCount`   
The number of secondary private IP addresses that Amazon Elastic Compute Cloud (Amazon EC2) automatically assigns to the network interface.

*Required*: No

*Type*: Integer

 `SubnetId`   
The ID of the subnet to associate with the network interface.

*Required*: Conditional. If you don't specify the `NetworkInterfaceId` property, you must specify this property.

*Type*: String


