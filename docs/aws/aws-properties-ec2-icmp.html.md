EC2 ICMP Property Type
======================

The EC2 ICMP property is an embedded property of the [AWS::EC2::NetworkAclEntry](aws-resource-ec2-network-acl-entry.html "AWS::EC2::NetworkAclEntry") type.

The following properties are available with the EC2 ICMP type.

Property

Type

Required

Notes

Code

Integer

Conditional

The Internet Control Message Protocol (ICMP) code. You can use -1 to specify all ICMP codes for the given ICMP type.

Condition: Required if specifying 1 (ICMP) for the CreateNetworkAclEntry protocol parameter.

Type

Integer

Conditional

The Internet Control Message Protocol (ICMP) type. You can use -1 to specify all ICMP types.

Condition: Required if specifying 1 (ICMP) for the CreateNetworkAclEntry protocol parameter.

