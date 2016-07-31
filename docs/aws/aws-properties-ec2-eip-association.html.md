AWS::EC2::EIPAssociation
========================

The AWS::EC2::EIPAssociation resource type associates an Elastic IP address with an Amazon EC2 instance. The Elastic IP address can be an existing Elastic IP address or an Elastic IP address allocated through an [AWS::EC2::EIP resource](aws-properties-ec2-eip.html "AWS::EC2::EIP").

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::EC2::EIPAssociation",
   "Properties": {
      "AllocationId": String,
      "EIP": String,
      "InstanceId": String,
      "NetworkInterfaceId": String,
      "PrivateIpAddress": String
   }
}
    
```

Properties
----------

 `AllocationId`   
Allocation ID for the VPC Elastic IP address you want to associate with an Amazon EC2 instance in your VPC.

*Required*: Conditional. Required for a VPC.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you also change the `InstanceId` or `NetworkInterfaceId` property. If not, update requires [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `EIP`   
Elastic IP address that you want to associate with the Amazon EC2 instance specified by the `InstanceId` property. You can specify an existing Elastic IP address or a reference to an Elastic IP address allocated with a [AWS::EC2::EIP resource](aws-properties-ec2-eip.html "AWS::EC2::EIP").

*Required*: Conditional. Required for Elastic IP addresses for use in EC2-Classic.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you also change the `InstanceId` or `NetworkInterfaceId` property. If not, update requires [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `InstanceId`   
Instance ID of the Amazon EC2 instance that you want to associate with the Elastic IP address specified by the EIP property.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you also change the `AllocationId` or `EIP` property. If not, update requires [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `NetworkInterfaceId`   
The ID of the network interface to associate with the Elastic IP address (VPC only).

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you also change the `AllocationId` or `EIP` property. If not, update requires [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `PrivateIpAddress`   
The private IP address that you want to associate with the Elastic IP address. The private IP address is restricted to the primary and secondary private IP addresses that are associated with the network interface. By default, the private IP address that is associated with the EIP is the primary private IP address of the network interface.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

For AWS::EC2::EIPAssociation snippets, see [Assigning an Amazon EC2 Elastic IP Using AWS::EC2::EIP Snippet](quickref-ec2.html#scenario-ec2-eip "Assigning an Amazon EC2 Elastic IP Using AWS::EC2::EIP Snippet").

