AWS::EC2::Subnet
================

Creates a subnet in an existing VPC.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::Subnet",
   "Properties" : {
      "AvailabilityZone" : String,
      "CidrBlock" : String,
      "MapPublicIpOnLaunch" : Boolean,
      "Tags" : [ Resource Tag, ... ],
      "VpcId" : { "Ref" : String }
   }
}     
    
```

Properties
----------

 `AvailabilityZone`   
The availability zone in which you want the subnet. Default: AWS selects a zone for you (recommended).

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

If you update this property, you must also update the `CidrBlock` property.

 `CidrBlock`   
The CIDR block that you want the subnet to cover (for example, `"10.0.0.0/24"`).

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

If you update this property, you must also update the `AvailabilityZone` property.

 `MapPublicIpOnLaunch`   
Indicates whether instances that are launched in this subnet receive a public IP address. By default, the value is `false`.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `Tags`   
An arbitrary set of tags (key–value pairs) for this subnet.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `VpcId`   
A Ref structure that contains the ID of the VPC on which you want to create the subnet. The VPC ID is provided as the value of the "Ref" property, as: `{                      "Ref": "VPCID" }`.

*Required*: Yes

*Type*: Ref ID

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

If you update this property, you must also update the `CidrBlock` property.

Return Values
-------------

You can pass the logical ID of the resource to an intrinsic function to get a value back from the resource. The value that is returned depends on the function used.

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource ID, such as `subnet-e19f0178`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `AvailabilityZone`   
Returns the availability zone (for example, `"us-east-1a"`) of this subnet.

Example:

``` {.programlisting}
              
{ "Fn::GetAtt" : [ "mySubnet", "AvailabilityZone" ] } 
            
```

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Example
-------

The following example snippet uses the VPC ID from a VPC named *myVPC* that was declared elsewhere in the same template.

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "mySubnet" : {
         "Type" : "AWS::EC2::Subnet",
         "Properties" : {
            "VpcId" : { "Ref" : "myVPC" },
            "CidrBlock" : "10.0.0.0/24",
            "AvailabilityZone" : "us-east-1a",
            "Tags" : [ { "Key" : "foo", "Value" : "bar" } ]
         }
      }
   }
}     
    
```

See Also
--------

-   [CreateSubnet](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSubnet.html) in the *Amazon EC2 API Reference*

-   [Using Tags](http://docs.aws.amazon.com/AWSEC2/latest/DeveloperGuide/Using_Tags.html) in the *Amazon Elastic Compute Cloud User Guide*


