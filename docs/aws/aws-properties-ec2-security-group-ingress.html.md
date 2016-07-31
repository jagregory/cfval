AWS::EC2::SecurityGroupIngress
==============================

The `AWS::EC2::SecurityGroupIngress` resource adds an ingress rule to an Amazon EC2 or Amazon VPC security group.

Important

Use `AWS::EC2::SecurityGroupIngress` and `AWS::EC2::SecurityGroupEgress` only when necessary, typically to allow security groups to reference each other in ingress and egress rules. Otherwise, use the embedded ingress and egress rules of [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup"). For more information, see [Amazon EC2 Security Groups](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).

Syntax
------

``` {.programlisting}
      {
   "CidrIp" : String,
   "FromPort" : Integer,
   "GroupId" : String,
   "GroupName" : String,
   "IpProtocol" : String,
   "SourceSecurityGroupName" : String,
   "SourceSecurityGroupId" : String,
   "SourceSecurityGroupOwnerId" : String,
   "ToPort" : Integer
}
    
```

Properties
----------

For more information about adding ingress rules to Amazon EC2 or VPC security groups, see [AuthorizeSecurityGroupIngress](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html) in the *Amazon EC2 API Reference*.

Note

If you change this resource's logical ID, you must also update a property value in order to trigger an update for this resource.

 `CidrIp`   
Specifies a CIDR range.

For an overview of CIDR ranges, go to the [Wikipedia Tutorial](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

*Type*: String

*Required*: Conditional. If you specify SourceSecurityGroupName, do not specify CidrIp.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `FromPort`   
Start of port range for the TCP and UDP protocols, or an ICMP type number. If you specify `icmp` for the `IpProtocol` property, you can specify -1 as a wildcard (i.e., any ICMP type number).

*Type*: Integer

*Required*: Yes, for ICMP and any protocol that uses ports.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `GroupId`   
ID of the Amazon EC2 or VPC security group to modify. The group must belong to your account.

*Type*: String

*Required*: Conditional. You must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property. For example, [EC2-VPC](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) accounts must use the GroupId property.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `GroupName`   
Name of the Amazon EC2 security group (non-VPC security group) to modify. This value can be a reference to an [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup") resource or the name of an existing Amazon EC2 security group.

*Type*: String

*Required*: Conditional. You must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property. For example, [EC2-VPC](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) accounts must use the GroupId property.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `IpProtocol`   
IP protocol name or number. For valid values, see the IpProtocol parameter in [AuthorizeSecurityGroupIngress](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html)

*Type*: String

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SourceSecurityGroupId`   
Specifies the ID of the source security group or uses the `Ref` intrinsic function to refer to the logical ID of a security group defined in the same template.

*Type*: String

*Required*: Conditional. If you specify CidrIp, do not specify SourceSecurityGroupId.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SourceSecurityGroupName`   
Specifies the name of the Amazon EC2 security group (non-VPC security group) to allow access or uses the `Ref` intrinsic function to refer to the logical name of a security group defined in the same template. For instances in a VPC, specify the `SourceSecurityGroupId` property.

*Type*: String

*Required*: Conditional. If you specify CidrIp, do not specify SourceSecurityGroupName.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SourceSecurityGroupOwnerId`   
Specifies the AWS Account ID of the owner of the Amazon EC2 security group specified in the SourceSecurityGroupName property.

*Type*: String

*Required*: Conditional. If you specify SourceSecurityGroupName and that security group is owned by a different account than the account creating the stack, you must specify the SourceSecurityGroupOwnerId; otherwise, this property is optional.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ToPort`   
End of port range for the TCP and UDP protocols, or an ICMP code. If you specify `icmp` for the `IpProtocol` property, you can specify -1 as a wildcard (i.e., any ICMP code).

*Type*: Integer

*Required*: Yes, for ICMP and any protocol that uses ports.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Examples
--------

EC2 Security Group and Ingress Rule

To create an Amazon EC2 (non-VPC) security group and an ingress rule, use the SourceSecurityGroupName property in the ingress rule.

The following template snippet creates an EC2 security group with an ingress rule that allows incoming traffic on port 80 from any other host in the security group. The snippet uses the intrinsic function [Ref](intrinsic-function-reference-ref.html "Ref") to specify the value for SourceSecurityGroupName.

``` {.programlisting}
      {
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "SGBase": {
            "Type": "AWS::EC2::SecurityGroup",
            "Properties": {
                "GroupDescription": "Base Security Group",
                "SecurityGroupIngress": [
                    {
                        "IpProtocol": "tcp",
                        "CidrIp": "0.0.0.0/0",
                        "FromPort": "22",
                        "ToPort": "22"
                    }
                ]
            }
        },
        "SGBaseIngress": {
            "Type": "AWS::EC2::SecurityGroupIngress",
            "Properties": {
                "GroupName": { "Ref": "SGBase" },
                "IpProtocol": "tcp",
                "FromPort": "80",
                "ToPort": "80",
                "SourceSecurityGroupName": { "Ref": "SGBase" }
            }
        }
    }
}           
    
```

VPC Security Groups with Egress and Ingress Rules

In some cases, you might have an originating (source) security group to which you want to add an outbound rule that allows traffic to a destination (target) security group. The target security group also needs an inbound rule that allows traffic from the source security group. Note that you cannot use the `Ref` function to specify the outbound and inbound rules for each security group. Doing so creates a circular dependency; you cannot have two resources that depend on each other. Instead, use the egress and ingress resources to declare these outbound and inbound rules, as shown in the following template snippet.

``` {.programlisting}
      {
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "SourceSG": {
      "Type": "AWS::EC2::SecurityGroup",
      "Properties": {
        "VpcId" : "vpc-e063f789",
        "GroupDescription": "Sample source security group"
      }
    },
    "TargetSG": {
      "Type": "AWS::EC2::SecurityGroup",
      "Properties": {
        "VpcId" : "vpc-e063f789",
        "GroupDescription": "Sample target security group"
      }
    },
    "OutboundRule": {
      "Type": "AWS::EC2::SecurityGroupEgress",
      "Properties":{
        "IpProtocol": "tcp",
        "FromPort": "0",
        "ToPort": "65535",
        "DestinationSecurityGroupId": {
          "Fn::GetAtt": [
            "TargetSG",
            "GroupId"
          ]
        },
        "GroupId": {
          "Fn::GetAtt": [
            "SourceSG",
            "GroupId"
          ]
        }
      }
    },
    "InboundRule": {
      "Type": "AWS::EC2::SecurityGroupIngress",
      "Properties":{
        "IpProtocol": "tcp",
        "FromPort": "0",
        "ToPort": "65535",
        "SourceSecurityGroupId": {
          "Fn::GetAtt": [
            "SourceSG",
            "GroupId"
          ]
        },
        "GroupId": {
          "Fn::GetAtt": [
            "TargetSG",
            "GroupId"
          ]
        }
      }
    }
  }
}
    
```

Allow Ping Requests

To allow ping requests, add the ICMP protocol type and specify `8` (echo request) for the ICMP type and either `0` or `-1` (all) for the ICMP code.

``` {.programlisting}
      "SGPing" : {
  "Type" : "AWS::EC2::SecurityGroup",
  "DependsOn": "VPC",
  "Properties" : {
    "GroupDescription" : "SG to test ping",
    "VpcId" : {"Ref" : "VPC"},
    "SecurityGroupIngress" : [ 
      { "IpProtocol" : "tcp", "FromPort" : "22", "ToPort" : "22", "CidrIp" : "10.0.0.0/24" },
      { "IpProtocol" : "icmp", "FromPort" : "8", "ToPort" : "-1", "CidrIp" : "10.0.0.0/24" }
    ]
  }
}
    
```
