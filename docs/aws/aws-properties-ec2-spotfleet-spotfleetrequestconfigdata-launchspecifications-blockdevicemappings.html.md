Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings
==========================================================================================================

`BlockDeviceMappings` is a property of the [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications") property that defines the block devices that are mapped to an instance.

Syntax
------

``` {.programlisting}
      {
  "DeviceName" : String,
  "Ebs" : EBSBlockDevice,
  "NoDevice" : Boolean,
  "VirtualName" : String
}
    
```

Properties
----------

 `DeviceName`   
The name of the device within the EC2 instance, such as `/dev/dsh` or `xvdh`.

*Required*: Yes

*Type*: String

 `Ebs`   
The Amazon Elastic Block Store (Amazon EBS) volume information.

*Required*: Conditional You can specify either the `VirtualName` or `Ebs`, but not both.

*Type*: [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings Ebs](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings Ebs")

 `NoDevice`   
Suppresses the specified device that is included in the block device mapping of the Amazon Machine Image (AMI).

*Required*: No

*Type*: Boolean

 `VirtualName`   
The name of the virtual device. The name must be in the form `ephemeralX` where *`X`* is a number equal to or greater than zero (0), for example, `ephemeral0`.

*Required*: Conditional You can specify either the `VirtualName` or `Ebs`, but not both.

*Type*: String


