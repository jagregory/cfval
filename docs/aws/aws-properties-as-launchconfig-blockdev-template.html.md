AWS CloudFormation AutoScaling EBS Block Device Property Type
=============================================================

The AutoScaling EBS Block Device type is an embedded property of the [AutoScaling Block Device Mapping](aws-properties-as-launchconfig-blockdev-mapping.html "AWS CloudFormation AutoScaling Block Device Mapping Property Type") type.

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
Indicates whether to delete the volume when the instance is terminated. By default, Auto Scaling uses `true`.

*Required*: No

*Type*: Boolean

 `Encrypted`   
Indicates whether the volume is encrypted. Encrypted EBS volumes must be attached to instances that support Amazon EBS encryption. Volumes that you create from encrypted snapshots are automatically encrypted. You cannot create an encrypted volume from an unencrypted snapshot or an unencrypted volume from an encrypted snapshot.

*Required*: No

*Type*: Boolean

 `Iops`   
The number of I/O operations per second (IOPS) that the volume supports. The maximum ratio of IOPS to volume size is 30.

*Required*: No

*Type*: Integer.

 `SnapshotId`   
The snapshot ID of the volume to use.

*Required*: Conditional If you specify both `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal or greater than the size of the snapshot.

*Type*: String

 `VolumeSize`   
The volume size, in Gibibytes (GiB). This can be a number from 1 – 1024. If the volume type is EBS optimized, the minimum value is 10. For more information about specifying the volume type, see EbsOptimized in [AWS::AutoScaling::LaunchConfiguration](aws-properties-as-launchconfig.html "AWS::AutoScaling::LaunchConfiguration").

*Required*: Conditional If you specify both `SnapshotId` and `VolumeSize`, `VolumeSize` must be equal or greater than the size of the snapshot.

*Type*: Integer.

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `VolumeType`   
The volume type. By default, Auto Scaling uses the `standard` volume type. For more information, see [Ebs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Ebs.html) in the *Auto Scaling API Reference*.

*Required*: No

*Type*: String

Examples
--------

For AutoScaling EBS Block Device snippets, see [Auto Scaling Launch Configuration Resource](quickref-autoscaling.html#scenario-as-launch-config "Auto Scaling Launch Configuration Resource").

