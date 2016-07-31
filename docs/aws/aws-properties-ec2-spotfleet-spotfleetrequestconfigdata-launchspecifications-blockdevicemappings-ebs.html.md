Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings Ebs
==============================================================================================================

`Ebs` is a property of the [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications BlockDeviceMappings") property that defines a block device for an Amazon Elastic Block Store (Amazon EBS) volume.

Syntax
------

``` {.programlisting}
      {
  "DeleteOnTermination" : Boolean,
  "Encrypted" : Boolean,
  "Iops" : Integer,
  "SnapshotId" : String,
  "VolumeSize" : Integer,
  "VolumeType" : String
}
    
```

Properties
----------

 `DeleteOnTermination`   
Indicates whether to delete the volume when the instance is terminated.

*Required*: No

*Type*: Boolean

 `Encrypted`   
Indicates whether the EBS volume is encrypted. Encrypted Amazon EBS volumes can be attached only to instances that support Amazon EBS encryption.

*Required*: No

*Type*: Boolean

 `Iops`   
The number of I/O operations per second (IOPS) that the volume supports. For more information, see [Iops](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) for the `EbsBlockDevice` action in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: Integer

 `SnapshotId`   
The snapshot ID of the volume that you want to use.

*Required*: Conditional If you specify both the `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal to or greater than the size of the snapshot.

*Type*: String

 `VolumeSize`   
The volume size, in Gibibytes (GiB). For more information about specifying the volume size, see [VolumeSize](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) for the `EbsBlockDevice` action in the *Amazon EC2 API Reference*.

*Required*: Conditional If you specify both the `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal to or greater than the size of the snapshot.

*Type*: Integer

 `VolumeType`   
The volume type. For more information about specifying the volume type, see [VolumeType](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) for the `EbsBlockDevice` action in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: String


