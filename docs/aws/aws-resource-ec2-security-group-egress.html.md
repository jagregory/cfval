AWS::EC2::SecurityGroupEgress
=============================

The `AWS::EC2::SecurityGroupEgress` resource adds an egress rule to an Amazon VPC security group.

Important

Use `AWS::EC2::SecurityGroupIngress` and `AWS::EC2::SecurityGroupEgress` only when necessary, typically to allow security groups to reference each other in ingress and egress rules. Otherwise, use the embedded ingress and egress rules of [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup"). For more information, see [Amazon EC2 Security Groups](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).

Syntax
------

``` {.programlisting}
      {
   "CidrIp" : String,
   "DestinationSecurityGroupId" : String,
   "FromPort" : Integer,
   "GroupId" : String,
   "IpProtocol" : String,
   "ToPort" : Integer
}
    
```

Properties
----------

For more information about adding egress rules to VPC security groups, go to [AuthorizeSecurityGroupEgress](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupEgress.html) in the *Amazon EC2 API Reference*.

Note

If you change this resource's logical ID, you must also update a property value in order to trigger an update for this resource.

 `CidrIp`   
CIDR range.

*Type*: String

*Required*: Conditional. Cannot be used when specifying a destination security group.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DestinationSecurityGroupId`   
Specifies the group ID of the destination Amazon VPC security group.

*Type*: String

*Required*: Conditional. Cannot be used when specifying a CIDR IP address.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `FromPort`   
Start of port range for the TCP and UDP protocols, or an ICMP type number. If you specify `icmp` for the `IpProtocol` property, you can specify -1 as a wildcard (i.e., any ICMP type number).

*Type*: Integer

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `GroupId`   
ID of the Amazon VPC security group to modify. This value can be a reference to an [AWS::EC2::SecurityGroup](aws-properties-ec2-security-group.html "AWS::EC2::SecurityGroup") resource that has a valid VpcId property or the ID of an existing Amazon VPC security group.

*Type*: String

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `IpProtocol`   
IP protocol name or number. For valid values, see the IpProtocol parameter in [AuthorizeSecurityGroupIngress](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html)

*Type*: String

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ToPort`   
End of port range for the TCP and UDP protocols, or an ICMP code. If you specify `icmp` for the `IpProtocol` property, you can specify -1 as a wildcard (i.e., any ICMP code).

*Type*: Integer

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

VPC Security Groups Example
---------------------------

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
