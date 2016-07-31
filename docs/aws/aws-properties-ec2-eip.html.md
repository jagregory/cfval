AWS::EC2::EIP
=============

The AWS::EC2::EIP resource allocates an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::EIP",
   "Properties" : {
      "InstanceId" : String,
      "Domain" : String
   }
}
    
```

Properties
----------

 `InstanceId`   
The Instance ID of the Amazon EC2 instance that you want to associate with this Elastic IP address.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Domain`   
Set to `vpc` to allocate the address to your Virtual Private Cloud (VPC). No other values are supported.

Note

If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the `DependsOn` attribute on this resource. For more information, see [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

For more information, see [AllocateAddress](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AllocateAddress.html) in the *Amazon EC2 API Reference*. For more information about Elastic IP Addresses in VPC, go to [IP Addressing in Your VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/vpc-ip-addressing.html) in the *Amazon VPC User Guide*.

*Required*: Conditional. Required when allocating an address to a VPC

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When you specify the logical ID of an AWS::EC2::EIP object as an argument to the `Ref` function, AWS CloudFormation returns the value of the instance's *`PublicIp`*.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `AllocationId`   
The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. Example return value: `eipalloc-5723d13e`

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

To view AWS::EC2::EIP snippets, see [Assigning an Amazon EC2 Elastic IP Using AWS::EC2::EIP Snippet](quickref-ec2.html#scenario-ec2-eip "Assigning an Amazon EC2 Elastic IP Using AWS::EC2::EIP Snippet").

