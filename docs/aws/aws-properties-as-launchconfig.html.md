AWS::AutoScaling::LaunchConfiguration
=====================================

The AWS::AutoScaling::LaunchConfiguration type creates an Auto Scaling launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances in the Auto Scaling group.

Important

When you update a property of the LaunchConfiguration resource, AWS CloudFormation deletes that resource and creates a new launch configuration with the updated properties and a new name. This update action does not deploy any change across the running Amazon EC2 instances in the auto scaling group. In other words, an update simply replaces the LaunchConfiguration so that when the auto scaling group launches new instances, they will get the updated configuration, but existing instances continue to run with the configuration that they were originally launched with. This works the same way as if you made similar changes manually to an auto scaling group.

If you want to update existing instances when you update the LaunchConfiguration resource, you must specify an update policy attribute for the `AWS::AutoScaling::AutoScalingGroup` resource. For more information, see [UpdatePolicy](aws-attribute-updatepolicy.html "UpdatePolicy Attribute").

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::AutoScaling::LaunchConfiguration",
   "Properties" : {
      "AssociatePublicIpAddress" : Boolean,
      "BlockDeviceMappings" : [ BlockDeviceMapping, ... ],
      "ClassicLinkVPCId" : String,
      "ClassicLinkVPCSecurityGroups" : [ String, ... ],
      "EbsOptimized" : Boolean,
      "IamInstanceProfile" : String,
      "ImageId" : String,
      "InstanceId" : String,
      "InstanceMonitoring" : Boolean,
      "InstanceType" : String,
      "KernelId" : String,
      "KeyName" : String,
      "PlacementTenancy" : String,
      "RamDiskId" : String,
      "SecurityGroups" : [ SecurityGroup, ... ],
      "SpotPrice" : String,
      "UserData" : String
   }
}     
    
```

Properties
----------

 `AssociatePublicIpAddress`   
For Amazon EC2 instances in a VPC, indicates whether instances in the Auto Scaling group receive public IP addresses. If you specify `true`, each instance in the Auto Scaling receives a unique public IP address.

Note

If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the `DependsOn` attribute to declare a dependency on the VPC-gateway attachment. For more information, see [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

*Required*: No

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `BlockDeviceMappings`   
Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.

*Required*: No

*Type*: A list of [BlockDeviceMappings](aws-properties-as-launchconfig-blockdev-mapping.html "AWS CloudFormation AutoScaling Block Device Mapping Property Type").

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ClassicLinkVPCId`   
The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to. You can specify this property only for EC2-Classic instances. For more information, see [ClassicLink](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the Amazon Elastic Compute Cloud User Guide.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ClassicLinkVPCSecurityGroups`   
The IDs of one or more security groups for the VPC that you specified in the `ClassicLinkVPCId` property.

*Required*: Conditional. If you specified the `ClassicLinkVPCId` property, you must specify this property.

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EbsOptimized`   
Specifies whether the launch configuration is optimized for EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal EBS I/O performance.

Additional fees are incurred when using EBS-optimized instances. For more information about fees and supported instance types, see [EBS-Optimized Instances](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No If this property is not specified, "false" is used.

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `IamInstanceProfile`   
Provides the name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role.

*Required*: No

*Type*: String (1–1600 chars)

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ImageId`   
Provides the unique ID of the Amazon Machine Image (AMI) that was assigned during registration.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `InstanceId`   
The ID of the Amazon EC2 instance you want to use to create the launch configuration. Use this property if you want the launch configuration to use settings from an existing Amazon EC2 instance.

When you use an instance to create a launch configuration, all properties are derived from the instance with the exception of `BlockDeviceMapping` and `AssociatePublicIpAddress`. You can override any properties from the instance by specifying them in the launch configuration.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `InstanceMonitoring`   
Indicates whether detailed instance monitoring is enabled for the Auto Scaling group. By default, this property is set to `true` (enabled).

When detailed monitoring is enabled, Amazon CloudWatch (CloudWatch) generates metrics every minute and your account is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes. For more information, see [Monitor Your Auto Scaling Instances](http://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html) in the *Auto Scaling Developer Guide*.

*Required*: No

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `InstanceType`   
Specifies the instance type of the EC2 instance.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `KernelId`   
Provides the ID of the kernel associated with the EC2 AMI.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `KeyName`   
Provides the name of the EC2 key pair.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `PlacementTenancy`   
The tenancy of the instance. An instance with a tenancy of `dedicated` runs on single-tenant hardware and can only be launched in a VPC. You must set the value of this parameter to `dedicated` if want to launch dedicated instances in a shared tenancy VPC (a VPC with the instance placement tenancy attribute set to default). For more information, see [CreateLaunchConfiguration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html) in the *Auto Scaling API Reference*.

If you specify this property, you must specify at least one subnet in the VPCZoneIdentifier property of the [AWS::AutoScaling::AutoScalingGroup](aws-properties-as-group.html "AWS::AutoScaling::AutoScalingGroup") resource.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `RamDiskId`   
The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, refer to the AWS Resource Center and search for the kernel ID.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SecurityGroups`   
A list that contains the EC2 security groups to assign to the Amazon EC2 instances in the Auto Scaling group. The list can contain the name of existing EC2 security groups or references to AWS::EC2::SecurityGroup resources created in the template. If your instances are launched within VPC, specify Amazon VPC security group IDs.

*Required*: No

*Type*: A list of EC2 security groups.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SpotPrice`   
The spot price for this autoscaling group. If a spot price is set, then the autoscaling group will launch when the current spot price is less than the amount specified in the template.

When you have specified a spot price for an auto scaling group, the group will only launch when the spot price has been met, regardless of the setting in the autoscaling group's *`DesiredCapacity`*.

For more information about configuring a spot price for an autoscaling group, see [Using Auto Scaling to Launch Spot Instances](http://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/US-SpotInstances.html) in the *AutoScaling Developer Guide*.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

When you change your bid price by creating a new launch configuration, running instances will continue to run as long as the bid price for those running instances is higher than the current Spot price.

 `UserData`   
The user data available to the launched EC2 instances.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "LaunchConfig" }
    
```

For the resource with the logical ID `LaunchConfig`, `Ref` will return the Auto Scaling launch configuration name, such as `mystack-mylaunchconfig-1DDYF1E3B3I`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Examples
-----------------

**Example LaunchConfig with block device**

This example shows a launch configuration that describes two Amazon Elastic Block Store mappings.

``` {.programlisting}
          "LaunchConfig" : {
   "Type" : "AWS::AutoScaling::LaunchConfiguration",
   "Properties" : {
      "KeyName" : { "Ref" : "KeyName" },
      "ImageId" : {
         "Fn::FindInMap" : [
            "AWSRegionArch2AMI",
            { "Ref" : "AWS::Region" },
            {
               "Fn::FindInMap" : [
                  "AWSInstanceType2Arch", { "Ref" : "InstanceType" }, "Arch"
               ]
            }
         ]
      },
      "UserData" : { "Fn::Base64" : { "Ref" : "WebServerPort" }},
      "SecurityGroups" : [ { "Ref" : "InstanceSecurityGroup" } ],
      "InstanceType" : { "Ref" : "InstanceType" },
      "BlockDeviceMappings" : [
         {
           "DeviceName" : "/dev/sda1",
           "Ebs" : { "VolumeSize" : "50", "VolumeType" : "io1", "Iops" : 200 } 
         },
         {
           "DeviceName" : "/dev/sdm",
           "Ebs" : { "VolumeSize" : "100", "DeleteOnTermination" : "true"}
         }
      ]
   }
} 
        
```

**Example LaunchConfig with Spot Price in Autoscaling Group**

This example shows a launch configuration that features a spot price in the AutoScaling group. This launch configuration will only be active if the current spot price is less than the amount in the template specification (0.05).

``` {.programlisting}
          
"LaunchConfig" : {
   "Type" : "AWS::AutoScaling::LaunchConfiguration",
   "Properties" : {
      "KeyName" : { "Ref" : "KeyName" },
      "ImageId" : {
         "Fn::FindInMap" : [
            "AWSRegionArch2AMI",
            { "Ref" : "AWS::Region" },
            {
               "Fn::FindInMap" : [
                  "AWSInstanceType2Arch", { "Ref" : "InstanceType" }, "Arch"
               ]
            }
         ]
      },
      "SecurityGroups" : [ { "Ref" : "InstanceSecurityGroup" } ],
      "SpotPrice" :  "0.05",
      "InstanceType" : { "Ref" : "InstanceType" }
   }
} 
        
```

**Example LaunchConfig with IAM Instance Profile**

Here's a launch configuration using the [IamInstanceProfile](aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile) property.

Only the AWS::AutoScaling::LaunchConfiguration specification is shown. For the full template, including the definition of, and further references from the [AWS::IAM::InstanceProfile](aws-resource-iam-instanceprofile.html "AWS::IAM::InstanceProfile") object referenced here as "RootInstanceProfile", see: [auto\_scaling\_with\_instance\_profile.template](https://s3.amazonaws.com/cloudformation-templates-us-east-1/auto_scaling_with_instance_profile.template).

``` {.programlisting}
          
"myLCOne": {
   "Type": "AWS::AutoScaling::LaunchConfiguration",
   "Properties": {
      "ImageId": {
         "Fn::FindInMap": [
            "AWSRegionArch2AMI",
            { "Ref": "AWS::Region" },
            {
               "Fn::FindInMap": [
                  "AWSInstanceType2Arch", { "Ref": "InstanceType" }, "Arch"
               ]
            }
         ]
      },
      "InstanceType": { "Ref": "InstanceType" },
      "IamInstanceProfile": { "Ref": "RootInstanceProfile" }
   }
}
        
```

**Example EBS-optimized volume with specified PIOPS**

You can create an AWS CloudFormation stack with auto scaled instances that contain EBS-optimized volumes with a specified PIOPS. This can increase the performance of your EBS-backed instances as explained in [Increasing EBS Performance](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSPerformance.html) in the *Amazon Elastic Compute Cloud User Guide*.

Caution

Additional fees are incurred when using EBS-optimized instances. For more information, see [EBS-Optimized Instances](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#EBSOptimized) in the *Amazon Elastic Compute Cloud User Guide*.

Because you cannot override PIOPS settings in an auto scaling launch configuration, the AMI in your launch configuration must have been configured with a block device mapping that specifies the desired PIOPS. You can do this by creating your own EC2 AMI with the following characteristics:

-   An instance type of `m1.large` or greater. This is required for EBS optimization.

-   An EBS-backed AMI with a volume type of "io1" and the number of IOPS you want for the Auto Scaling-launched instances.

-   The size of the EBS volume must accommodate the IOPS you need. There is a 10 : 1 ratio between IOPS and Gibibytes (GiB) of storage, so for 100 PIOPS, you need at least 10 GiB storage on the root volume.

Use this AMI in your Auto Scaling launch configuration. For example, an EBS-optimized AMI with PIOPS that has the AMI ID `ami-7430ba44` would be used in your launch configuration like this:

``` {.programlisting}
          
"LaunchConfig" : {
   "Type" : "AWS::AutoScaling::LaunchConfiguration",
   "Properties" : {
      "KeyName" : { "Ref" : "KeyName" },
      "ImageId" : "ami-7430ba44",
      "UserData" : { "Fn::Base64" : { "Ref" : "WebServerPort" } },
      "SecurityGroups" : [ { "Ref" : "InstanceSecurityGroup" } ],
      "InstanceType" : "m1.large",
      "EbsOptimized" : "true"
   }
},       
        
```

Be sure to set the *`InstanceType`* to at least *m1.large* and set *`EbsOptimized`* to *true*.

When you create a launch configuration such as this one, your launched instances will contain optimized EBS root volumes with the PIOPS that you selected when creating the AMI.

To view more LaunchConfiguration snippets, see [Auto Scaling Launch Configuration Resource](quickref-autoscaling.html#scenario-as-launch-config "Auto Scaling Launch Configuration Resource").

See Also
--------

-   [Creating Your Own AMIs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html) in the *Amazon Elastic Compute Cloud User Guide*.

-   [Block Device Mapping](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the *Amazon Elastic Compute Cloud User Guide*.


