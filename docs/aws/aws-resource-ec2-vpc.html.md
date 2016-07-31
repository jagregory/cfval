AWS::EC2::VPC
=============

Creates a Virtual Private Cloud (VPC) with the CIDR block that you specify. To name a VPC resource, use the `Tags` property and specify a value for the `Name` key.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::VPC",
   "Properties" : {
      "CidrBlock" : String,
      "EnableDnsSupport" : Boolean,
      "EnableDnsHostnames" : Boolean,
      "InstanceTenancy" : String,
      "Tags" : [ Resource Tag, ... ]
   }
}     
    
```

Properties
----------

 `CidrBlock`   
The CIDR block you want the VPC to cover. For example: `"10.0.0.0/16"`.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EnableDnsSupport`   
Specifies whether DNS resolution is supported for the VPC. If this attribute is `true`, the Amazon DNS server resolves DNS hostnames for your instances to their corresponding IP addresses; otherwise, it does not. By default the value is set to `true`.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EnableDnsHostnames`   
Specifies whether the instances launched in the VPC get DNS hostnames. If this attribute is `true`, instances in the VPC get DNS hostnames; otherwise, they do not. You can only set `EnableDnsHostnames` to `true` if you also set the `EnableDnsSupport` attribute to `true`. By default, the value is set to `false`.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `InstanceTenancy`   
The allowed tenancy of instances launched into the VPC.

-   `"default"`: Instances can be launched with any tenancy.

-   `"dedicated"`: Any instance launched into the VPC automatically has dedicated tenancy, unless you launch it with the default tenancy.

*Required*: No

*Type*: String

*Valid values*: `"default"` or `"dedicated"`

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this VPC. To name a VPC resource, specify a value for the `Name` key.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource ID, such as `vpc-18ac277d`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

You can obtain the following default resource IDs, which AWS creates whenever you create a VPC.

 `CidrBlock`   
The set of IP addresses for the VPC. For example, `10.0.0.0/16`.

 `DefaultNetworkAcl`   
The default network ACL ID that is associated with the VPC. For example, `acl-814dafe3`.

 `DefaultSecurityGroup`   
The default security group ID that is associated with the VPC. For example, `sg-b178e0d3`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myVPC" : {
         "Type" : "AWS::EC2::VPC",
         "Properties" : {
            "CidrBlock" : "10.0.0.0/16",
            "EnableDnsSupport" : "false",
            "EnableDnsHostnames" : "false",
            "InstanceTenancy" : "dedicated",
            "Tags" : [ {"Key" : "foo", "Value" : "bar"} ]
         }
      }
   }
}     
    
```

See Also
--------

-   [CreateVpc](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpc.html) in the *Amazon EC2 API Reference*.


