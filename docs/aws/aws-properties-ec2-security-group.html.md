AWS::EC2::SecurityGroup
=======================

Creates an Amazon EC2 security group. To create a VPC security group, use the [VpcId](aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-vpcid) property.

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Important

If you want to cross-reference two security groups in the ingress and egress rules of those security groups, use the [AWS::EC2::SecurityGroupEgress](aws-resource-ec2-security-group-egress.html "AWS::EC2::SecurityGroupEgress") and [AWS::EC2::SecurityGroupIngress](aws-properties-ec2-security-group-ingress.html "AWS::EC2::SecurityGroupIngress") resources to define your rules. Do not use the embedded ingress and egress rules in the `AWS::EC2::SecurityGroup`. If you do, it causes a circular dependency, which AWS CloudFormation doesn't allow.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EC2::SecurityGroup",
  "Properties" : {
     "GroupDescription" : String,
     "SecurityGroupEgress" : [ Security Group Rule, ... ],
     "SecurityGroupIngress" : [ Security Group Rule, ... ],
     "Tags" :  [ Resource Tag, ... ],
     "VpcId" : String
  }
}
    
```

Properties
----------

 `GroupDescription`   
Description of the security group.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SecurityGroupEgress`   
A list of Amazon EC2 security group egress rules.

*Required*: No

*Type*: List of [EC2 Security Group Rule](aws-properties-ec2-security-group-rule.html "EC2 Security Group Rule Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SecurityGroupIngress`   
A list of Amazon EC2 security group ingress rules.

*Required*: No

*Type*: List of [EC2 Security Group Rule](aws-properties-ec2-security-group-rule.html "EC2 Security Group Rule Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Tags`   
The tags that you want to attach to the resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `VpcId`   
The physical ID of the VPC. Can be obtained by using a reference to an [AWS::EC2::VPC](aws-resource-ec2-vpc.html "AWS::EC2::VPC"), such as: `{ "Ref" : "myVPC" }`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

*Required*: Yes, for VPC security groups

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Note

For more information about VPC security groups, go to [Security Groups](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the *Amazon VPC User Guide*.

Return Values
-------------

### Ref

When you specify an AWS::EC2::SecurityGroup type as an argument to the `Ref` function, AWS CloudFormation returns the security group name or the security group ID (for EC2-VPC security groups that are not in a default VPC).

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `GroupId`   
The group ID of the specified security group, such as `sg-94b3a1f6`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

The following sample defines a security group with an ingress and egress rule:

``` {.programlisting}
      "InstanceSecurityGroup" : {
   "Type" : "AWS::EC2::SecurityGroup",
   "Properties" : {
      "GroupDescription" : "Allow http to client host",
      "VpcId" : {"Ref" : "myVPC"},
      "SecurityGroupIngress" : [{
            "IpProtocol" : "tcp",
            "FromPort" : "80",
            "ToPort" : "80",
            "CidrIp" : "0.0.0.0/0"
         }],
      "SecurityGroupEgress" : [{
         "IpProtocol" : "tcp",
         "FromPort" : "80",
         "ToPort" : "80",
         "CidrIp" : "0.0.0.0/0"
      }]
   }
}     
    
```

When you create a VPC security group, Amazon EC2 creates a default egress rule that allows egress traffic on all ports and IP protocols to any location. The default rule is removed only when you specify one or more egress rules. If you want to remove the default rule and limit egress traffic to just the localhost (`127.0.0.1/32`), you can use the following sample:

``` {.programlisting}
      "sgwithoutegress": {
  "Type": "AWS::EC2::SecurityGroup",
  "Properties": {
    "GroupDescription": "Limits security group egress traffic",
    "SecurityGroupEgress": [
      {
        "CidrIp": "127.0.0.1/32",
        "IpProtocol": "-1"
      }
    ],
    "VpcId": { "Ref": "myVPC"}
  }
}
    
```

See Also
--------

-   [Using Security Groups](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html) in the *Amazon EC2 User Guide for Linux Instances*.

-   [Security Groups](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the *Amazon VPC User Guide*.


