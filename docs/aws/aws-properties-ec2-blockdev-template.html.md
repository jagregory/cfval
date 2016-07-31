Amazon Elastic Block Store Block Device Property
================================================

The Amazon Elastic Block Store block device type is an embedded property of the [Amazon EC2 Block Device Mapping Property](aws-properties-ec2-blockdev-mapping.html "Amazon EC2 Block Device Mapping Property") property.

Syntax
------

``` {.programlisting}
      
{
   "DeleteOnTermination" : Boolean,
   "Encrypted" : Boolean,
   "Iops" : Number,
   "SnapshotId" : String,
   "VolumeSize" : String,
   "VolumeType" : String
}     
    
```

Properties
----------

 `DeleteOnTermination`   
Determines whether to delete the volume on instance termination. The default value is `true`.

*Required*: No

*Type*: Boolean

 `Encrypted`   
Indicates whether the volume is encrypted. Encrypted Amazon EBS volumes can only be attached to instance types that support Amazon EBS encryption. Volumes that are created from encrypted snapshots are automatically encrypted. You cannot create an encrypted volume from an unencrypted snapshot or vice versa. If your AMI uses encrypted volumes, you can only launch the AMI on supported instance types. For more information, see [Amazon EBS encryption](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: Boolean

 `Iops`   
The number of I/O operations per second (IOPS) that the volume supports. This can be an integer from 100 – 2000.

*Required*: Conditional Required when the [volume type](aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-volumetype) is `io1`; not used with other volume types.

*Type*: Number

 `SnapshotId`   
The snapshot ID of the volume to use to create a block device.

*Required*: Conditional If you specify both `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal or greater than the size of the snapshot.

*Type*: String

 `VolumeSize`   
The volume size, in gibibytes (GiB). This can be a number from 1 – 1024. If the volume type is `io1`, the minimum value is 10.

*Required*: Conditional If you specify both `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal or greater than the size of the snapshot.

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `VolumeType`   
The volume type. If you set the type to `io1`, you must also set the Iops property. For valid values, see the `VolumeType` parameter for the [CreateVolume](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) action in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: String

Example
-------

``` {.programlisting}
      
{
   "DeviceName":"/dev/sdc",
   "Ebs":{
      "SnapshotId":"snap-xxxxxx",
      "VolumeSize":"50",
      "VolumeType":"io1",
      "Iops":"1000",
      "DeleteOnTermination":"false"
   }
}     
    
```

See Also
--------

-   [CreateVolume](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) in the *Amazon Elastic Compute Cloud API Reference*


