Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications
======================================================================================

`LaunchSpecifications` is a property of the [Amazon EC2 SpotFleet SpotFleetRequestConfigData](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html "Amazon EC2 SpotFleet SpotFleetRequestConfigData") property that defines the launch specifications for the Spot fleet request.

Syntax
------

``` {.programlisting}
      {
  "BlockDeviceMappings" : [ BlockDeviceMapping, ... ],
  "EbsOptimized" : Boolean,
  "IamInstanceProfile" : IamInstanceProfile,
  "ImageId" : String,
  "InstanceType" : String,
  "KernelId" : String,
  "KeyName" : String,
  "Monitoring" : Boolean,
  "NetworkInterfaces" : [ NetworkInterface, ... ],
  "Placement" : Placement,
  "RamdiskId" : String,
  "SecurityGroups" : [ SecurityGroup, ... ],
  "SubnetId" : String,
  "UserData" : String,
  "WeightedCapacity" : Number
}
    
```

Properties
----------

 `BlockDeviceMappings`   
Defines the block devices that are mapped to the Spot instances.

*Required*: No

*Type*: List of [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings")

 `EbsOptimized`   
Indicates whether the instances are optimized for Amazon Elastic Block Store (Amazon EBS) I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when you use an Amazon EBS-optimized instance.

*Required*: No

*Type*: Boolean

 `IamInstanceProfile`   
Defines the AWS Identity and Access Management (IAM) instance profile to associate with the instances.

*Required*: No

*Type*: [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications IamInstanceProfile](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications IamInstanceProfile")

 `ImageId`   
The unique ID of the Amazon Machine Image (AMI) to launch on the instances.

*Required*: Yes

*Type*: String

 `InstanceType`   
Specifies the instance type of the EC2 instances.

*Required*: Yes

*Type*: String

 `KernelId`   
The ID of the kernel that is associated with the Amazon Elastic Compute Cloud (Amazon EC2) AMI.

*Required*: No

*Type*: String

 `KeyName`   
An Amazon EC2 key pair to associate with the instances.

*Required*: No

*Type*: String

 `Monitoring`   
Enable or disable monitoring for the instances.

*Required*: No

*Type*: [Amazon EC2 SpotFleet SpotFleetRequestConfigData LaunchSpecifications Monitoring](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html "Amazon EC2 SpotFleet SpotFleetRequestConfigData LaunchSpecifications Monitoring")

 `NetworkInterfaces`   
The network interfaces to associate with the instances.

*Required*: No

*Type*: List of [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications NetworkInterfaces")

 `Placement`   
Defines a placement group, which is a logical grouping of instances within a single Availability Zone (AZ).

*Required*: No

*Type*: [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications Placement](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications Placement")

 `RamdiskId`   
The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, refer to the AWS Resource Center and search for the kernel ID.

*Required*: No

*Type*: String

 `SecurityGroups`   
One or more security group IDs to associate with the instances.

*Required*: No

*Type*: List of [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications SecurityGroups](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications SecurityGroups")

 `SubnetId`   
The ID of the subnet in which to launch the instances.

*Required*: No

*Type*: String

 `UserData`   
Base64-encoded MIME user data that instances use when starting up.

*Required*: No

*Type*: String

 `WeightedCapacity`   
The number of units provided by the specified instance type. These units are the same units that you chose to set the target capacity in terms of instances or a performance characteristic, such as vCPUs, memory, or I/O. For more information, see [How Spot Fleet Works](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances*.

If the target capacity divided by this value is not a whole number, Amazon EC2 rounds the number of instances to the next whole number.

*Required*: No

*Type*: Number


