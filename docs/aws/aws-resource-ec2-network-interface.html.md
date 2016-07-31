AWS::EC2::NetworkInterface
==========================

Describes a network interface in an Elastic Compute Cloud (EC2) instance for AWS CloudFormation. This is provided in a list in the NetworkInterfaces property of [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance").

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::NetworkInterface",
   "Properties" : {
      "Description" : String,
      "GroupSet" : [ String, ... ],
      "PrivateIpAddress" : String,
      "PrivateIpAddresses" : [ PrivateIpAddressSpecification, ... ],
      "SecondaryPrivateIpAddressCount" : Integer,
      "SourceDestCheck" : Boolean,
      "SubnetId" : String,
      "Tags" : [ Resource Tag, ... ]
   }
}
      
    
```

Properties
----------

 `Description`   
The description of this network interface.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `GroupSet`   
A list of security group IDs associated with this network interface.

*Required*: No

*Type*: List of strings.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PrivateIpAddress`   
Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the `PrivateIpAddresses` property.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `PrivateIpAddresses`   
Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the `Primary` property to `true` in the `PrivateIpAddressSpecification` property. If you want Amazon EC2 to automatically assign private IP addresses, use the `SecondaryPrivateIpAddressCount` property and do not specify this property.

For information about the maximum number of private IP addresses, see [Private IP Addresses Per ENI Per Instance Type](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: list of [PrivateIpAddressSpecification](aws-properties-ec2-network-interface-privateipspec.html "EC2 Network Interface Private IP Specification").

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you change the primary private IP address. If not, update requires [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `SecondaryPrivateIpAddressCount`   
The number of secondary private IP addresses that Amazon EC2 automatically assigns to the network interface. Amazon EC2 uses the value of the `PrivateIpAddress` property as the primary private IP address. If you don't specify that property, Amazon EC2 automatically assigns both the primary and secondary private IP addresses.

If you want to specify your own list of private IP addresses, use the `PrivateIpAddresses` property and do not specify this property.

For information about the maximum number of private IP addresses, see [Private IP Addresses Per ENI Per Instance Type](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: No

*Type*: Integer.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `SourceDestCheck`   
Flag indicating whether traffic to or from the instance is validated.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `SubnetId`   
The ID of the subnet to associate with the network interface.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `Tags`   
An arbitrary set of tags (key–value pairs) for this network interface.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `PrimaryPrivateIpAddress`   
Returns the primary private IP address of the network interface. For example, `10.0.0.192`.

 `SecondaryPrivateIpAddresses`   
Returns the secondary private IP addresses of the network interface. For example, `["10.0.0.161", "10.0.0.162", "10.0.0.163"]`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

Tip

For more NetworkInterface template examples, see [Elastic Network Interface (ENI) Template Snippets](quickref-ec2.html#cfn-template-snippets-eni "Elastic Network Interface (ENI) Template Snippets").

### Simple Standalone ENI

This is a simple standalone Elastic Network Interface (ENI), using all of the available properties.

``` {.programlisting}
        
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Description" : "Simple Standalone ENI",
   "Resources" : {
      "myENI" : {
         "Type" : "AWS::EC2::NetworkInterface",
         "Properties" : {
            "Tags": [{"Key":"foo","Value":"bar"}],
            "Description": "A nice description.",
            "SourceDestCheck": "false",
            "GroupSet": ["sg-75zzz219"],
            "SubnetId": "subnet-3z648z53",
            "PrivateIpAddress": "10.0.0.16"
         }
      }
   }
}        
      
```

### ENI on an EC2 instance

This is an example of an ENI on an EC2 instance. In this example, one ENI is added to the instance. If you want to add more than one ENI, you can specify a list for the `NetworkInterface` property. However, you can specify multiple ENIs only if all the ENIs have just private IP addresses (no associated public IP address). If you have an ENI with a public IP address, specify it and then use the `AWS::EC2::NetworkInterfaceAttachment` resource to add additional ENIs.

``` {.programlisting}
        
"Ec2Instance" : {
   "Type" : "AWS::EC2::Instance",
   "Properties" : {
      "ImageId" : { "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "AMI" ]},
      "KeyName" : { "Ref" : "KeyName" },
      "SecurityGroupIds" : [{ "Ref" : "WebSecurityGroup" }],
      "SubnetId" : { "Ref" : "SubnetId" },
      "NetworkInterfaces" : [ {
         "NetworkInterfaceId" : {"Ref" : "controlXface"}, "DeviceIndex" : "1" } ],
      "Tags" : [ {"Key" : "Role", "Value" : "Test Instance"}],
      "UserData" : { "Fn::Base64" : { "Ref" : "WebServerPort" }}
   }
}        
      
```

See Also
--------

-   [NetworkInterfaceType](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkInterfaceType.html) in the *Amazon Elastic Compute Cloud API Reference*


