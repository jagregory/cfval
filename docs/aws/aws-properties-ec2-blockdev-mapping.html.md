Amazon EC2 Block Device Mapping Property
========================================

The Amazon EC2 block device mapping property is an embedded property of the [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance") resource. For block device mappings for an Auto Scaling launch configuration, see [AutoScaling Block Device Mapping](aws-properties-as-launchconfig-blockdev-mapping.html "AWS CloudFormation AutoScaling Block Device Mapping Property Type").

Syntax
------

``` {.programlisting}
      
{
   "DeviceName" : String,
   "Ebs" : EC2 EBS Block Device,
   "NoDevice" : {},
   "VirtualName" : String
}     
    
```

Properties
----------

 `DeviceName`   
The name of the device within Amazon EC2.

*Required*: Yes

*Type*: String

 `Ebs`   
*Required*: Conditional You can specify either `VirtualName` or `Ebs`, but not both.

*Type*: [Amazon Elastic Block Store Block Device Property](aws-properties-ec2-blockdev-template.html "Amazon Elastic Block Store Block Device Property").

 `NoDevice`   
This property can be used to unmap a defined device.

*Required*: No

*Type*: an empty map: {}.

 `VirtualName`   
The name of the virtual device. The name must be in the form `ephemeralX` where *`X`* is a number starting from zero (0); for example, `ephemeral0`.

*Required*: Conditional You can specify either `VirtualName` or `Ebs`, but not both.

*Type*: String

Examples
--------

### Block Device Mapping with two EBS Volumes

This example sets the EBS-backed root device (/dev/sda1) size to 50 GiB, and another EBS-backed device mapped to /dev/sdm that is 100 GiB in size.

``` {.programlisting}
        
"BlockDeviceMappings" : [
   {
      "DeviceName" : "/dev/sda1",
      "Ebs" : { "VolumeSize" : "50" }
   },
   {
      "DeviceName" : "/dev/sdm",
      "Ebs" : { "VolumeSize" : "100" }
   }
]        
      
```

### Block Device Mapping with an Ephemeral Drive

This example maps an ephemeral drive to device /dev/sdc.

``` {.programlisting}
        
"BlockDeviceMappings" : [
   {
      "DeviceName"  : "/dev/sdc",
      "VirtualName" : "ephemeral0"
   }
]        
      
```

### Unmapping an AMI-defined Device

To unmap a device defined in the AMI, set the NoDevice property to an empty map, as shown here:

``` {.programlisting}
        
{
   "DeviceName":"/dev/sde",
   "NoDevice": {}
}        
      
```

See Also
--------

-   [Amazon EC2 Instance Store](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon Elastic Compute Cloud User Guide*


