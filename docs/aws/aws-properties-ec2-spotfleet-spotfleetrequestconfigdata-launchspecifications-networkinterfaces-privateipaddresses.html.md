Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces PrivateIpAddresses
===========================================================================================================================

`PrivateIpAddresses` is a property of the [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces") property that specifies the private IP address that you want to assign to the network interface.

Syntax
------

``` {.programlisting}
      {
  "Primary" : Boolean,
  "PrivateIpAddress" : String
}
    
```

Properties
----------

 `Primary`   
Indicates whether the private IP address is the primary private IP address. You can designate only one IP address as primary.

*Required*: No

*Type*: Boolean

 `PrivateIpAddress`   
The private IP address.

*Required*: Yes

*Type*: String


