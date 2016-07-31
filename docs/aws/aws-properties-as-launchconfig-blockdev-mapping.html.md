AWS CloudFormation AutoScaling Block Device Mapping Property Type
=================================================================

The AutoScaling Block Device Mapping type is an embedded property of the [AWS::AutoScaling::LaunchConfiguration](aws-properties-as-launchconfig.html "AWS::AutoScaling::LaunchConfiguration") type.

Syntax
------

``` {.programlisting}
      
{
   "DeviceName" : String,
   "Ebs" : AutoScaling EBS Block Device,
   "NoDevice" : Boolean,
   "VirtualName" : String
}     
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see [Ebs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Ebs.html) in the *Auto Scaling API Reference*.

 `DeviceName`   
The name of the device within Amazon EC2.

*Required*: Yes

*Type*: String

 `Ebs`   
The Amazon Elastic Block Store volume information.

*Required*: Conditional You can specify either `VirtualName` or `Ebs`, but not both.

*Type*: [AutoScaling EBS Block Device](aws-properties-as-launchconfig-blockdev-template.html "AWS CloudFormation AutoScaling EBS Block Device Property Type").

 `NoDevice`   
Suppresses the device mapping. If `NoDevice` is set to true for the root device, the instance might fail the Amazon EC2 health check. Auto Scaling launches a replacement instance if the instance fails the health check.

*Required*: No

*Type*: Boolean

 `VirtualName`   
The name of the virtual device. The name must be in the form `ephemeralX` where *`X`* is a number starting from zero (0), for example, `ephemeral0`.

*Required*: Conditional You can specify either `VirtualName` or `Ebs`, but not both.

*Type*: String


