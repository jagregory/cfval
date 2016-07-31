EC2 MountPoint Property Type
============================

The EC2 MountPoint property is an embedded property of the [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance") type.

Syntax
------

``` {.programlisting}
      
{
   "Device" : String,
   "VolumeId" : String
}     
    
```

Properties
----------

 `Device`   
How the device is exposed to the instance (such as /dev/sdh, or xvdh).

*Required*: Yes

*Type*: String

 `VolumeId`   
The ID of the Amazon EBS volume. The volume and instance must be within the same Availability Zone and the instance must be running.

*Required*: Yes

*Type*: String

Example
-------

This mount point (specified in the *`Volumes`* property in the EC2 instance) refers to a named EBS volume, "NewVolume".

``` {.programlisting}
      
"Ec2Instance" : {
   "Type" : "AWS::EC2::Instance",
   "Properties" : {
      "AvailabilityZone" : {
         "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "TestAz" ]
      },
      "SecurityGroups" : [ { "Ref" : "InstanceSecurityGroup" } ],
      "KeyName" : { "Ref" : "KeyName" },
      "ImageId" : {
         "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "AMI" ]
      },
      "Volumes" : [
         { "VolumeId" : { "Ref" : "NewVolume" }, "Device" : "/dev/sdk" }
      ]
   }
},
"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
      "Size" : "100",
      "AvailabilityZone" : {
         "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "TestAz" ]
      }
   }
}

    
```

See Also
--------

-   [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance")
-   [AWS::EC2::Volume](aws-properties-ec2-ebs-volume.html "AWS::EC2::Volume")

