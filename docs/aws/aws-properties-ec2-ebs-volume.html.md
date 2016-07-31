AWS::EC2::Volume
================

The AWS::EC2::Volume type creates a new Amazon Elastic Block Store (Amazon EBS) volume.

You can set a deletion policy for your volume to control how AWS CloudFormation handles the volume when the stack is deleted. For Amazon EBS volumes, you can choose to *retain* the volume, to *delete* the volume, or to *create a snapshot* of the volume. For more information, see [DeletionPolicy Attribute](aws-attribute-deletionpolicy.html "DeletionPolicy Attribute").

Note

If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.

Syntax
------

``` {.programlisting}
      
{
   "Type":"AWS::EC2::Volume",
   "Properties" : {
      "AutoEnableIO" : Boolean,
      "AvailabilityZone" : String,
      "Encrypted" : Boolean,
      "Iops" : Number,
      "KmsKeyId" : String,
      "Size" : String,
      "SnapshotId" : String,
      "Tags" : [ Resource Tag, ... ],
      "VolumeType" : String
   }
}
    
```

Properties
----------

 `AutoEnableIO`   
Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O. For more information, see [Working with the AutoEnableIO Volume Attribute](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-status.html#volumeIO) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AvailabilityZone`   
The Availability Zone in which to create the new volume.

*Required*: Yes

*Type*: String

*Update requires*: Updates are not supported.

 `Encrypted`   
Indicates whether the volume is encrypted. Encrypted Amazon EBS volumes can only be attached to instance types that support Amazon EBS encryption. Volumes that are created from encrypted snapshots are automatically encrypted. You cannot create an encrypted volume from an unencrypted snapshot or vice versa. If your AMI uses encrypted volumes, you can only launch the AMI on supported instance types. For more information, see [Amazon EBS encryption](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: Conditional. If you specify the `KmsKeyId` property, you must enable encryption.

*Type*: Boolean

*Update requires*: Updates are not supported.

 `Iops`   
The number of I/O operations per second (IOPS) that the volume supports. For more information about the valid sizes for each volume type, see the `Iops` parameter for the [`CreateVolume`](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) action in the *Amazon EC2 API Reference*.

*Required*: Conditional. *Required* when the volume type is `io1`; not used with other volume types.

*Type*: Number

*Update requires*: Updates are not supported.

 `KmsKeyId`   
The Amazon Resource Name (ARN) of the AWS Key Management Service master key that is used to create the encrypted volume, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`. If you create an encrypted volume and don't specify this property, the default master key is used.

*Required*: No

*Type*: String

*Update requires*: Updates are not supported.

 `Size`   
The size of the volume, in gibibytes (GiBs). For more information about the valid sizes for each volume type, see the `Size` parameter for the [`CreateVolume`](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) action in the *Amazon EC2 API Reference*.

If you specify the `SnapshotId` property, specify a size that is equal to or greater than the snapshot size. If you don't specify a size, Amazon EC2 will use the size of the snapshot as the volume size.

*Required*: Conditional. If you don't specify a value for the `SnapshotId` property, you must specify this property.

*Type*: String

*Update requires*: Updates are not supported.

 `SnapshotId`   
The snapshot from which to create the new volume.

*Required*: No

*Type*: String

*Update requires*: Updates are not supported.

 `Tags`   
An arbitrary set of tags (key–value pairs) for this volume.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VolumeType`   
The volume type. If you set the type to `io1`, you must also set the Iops property. For valid values, see the `VolumeType` parameter for the [CreateVolume](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) action in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: String

*Update requires*: Updates are not supported.

Return Values
-------------

### Ref

When you specify an AWS::EC2::Volume type as an argument to the `Ref` function, AWS CloudFormation returns the volume's physical ID. For example: `vol-5cb85026`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

**Example Encrypted Amazon EBS volume with DeletionPolicy to make a snapshot on delete**

``` {.programlisting}
          
"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
      "Size" : "100",
      "Encrypted" : "true",
      "AvailabilityZone" : { "Fn::GetAtt" : [ "Ec2Instance", "AvailabilityZone" ] },
      "Tags" : [ {
         "Key" : "MyTag",
         "Value" : "TagValue"
      } ]
   },
   "DeletionPolicy" : "Snapshot"
}
         
        
```

**Example Amazon EBS volume with 100 provisioned IOPS**

``` {.programlisting}
          
"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
     "Size" : "100",
     "VolumeType" : "io1",
     "Iops" : "100",
     "AvailabilityZone" : { "Fn::GetAtt" : [ "EC2Instance", "AvailabilityZone" ] }
   }
}        
        
```

See Also
--------

-   [CreateVolume](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html) in the *Amazon Elastic Compute Cloud API Reference*

-   [DeletionPolicy Attribute](aws-attribute-deletionpolicy.html "DeletionPolicy Attribute")


