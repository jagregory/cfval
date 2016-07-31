EC2 PortRange Property Type
===========================

The EC2 PortRange property is an embedded property of the [AWS::EC2::NetworkAclEntry](aws-resource-ec2-network-acl-entry.html "AWS::EC2::NetworkAclEntry") type.

The following properties are available with the EC2 PortRange type.

Property

Type

Required

Notes

From

Integer

Conditional

The first port in the range.

Condition: Required if specifying 6 (TCP) or 17 (UDP) for the CreateNetworkAclEntry protocol parameter.

To

Integer

Conditional

The last port in the range.

Condition: Required if specifying 6 (TCP) or 17 (UDP) for the CreateNetworkAclEntry protocol parameter.

