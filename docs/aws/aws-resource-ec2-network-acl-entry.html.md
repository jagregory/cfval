AWS::EC2::NetworkAclEntry
=========================

Creates an entry (i.e., rule) in a network ACL with a rule number you specify. Each network ACL has a set of numbered ingress rules and a separate set of numbered egress rules.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::NetworkAclEntry",
   "Properties" : {
      "CidrBlock" : String,
      "Egress" : Boolean,
      "Icmp" : EC2 ICMP,
      "NetworkAclId" : String,
      "PortRange" : EC2 PortRange,
      "Protocol" : Integer,
      "RuleAction" : String,
      "RuleNumber" : Integer
   }
}     
    
```

Properties
----------

 `CidrBlock`   
The CIDR range to allow or deny, in CIDR notation (e.g., 172.16.0.0/24).

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Egress`   
Whether this rule applies to egress traffic from the subnet (`true`) or ingress traffic to the subnet (`false`). By default, AWS CloudFormation specifies `false`.

*Required*: No

*Type*: Boolean

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `Icmp`   
The Internet Control Message Protocol (ICMP) code and type.

*Required*: Conditional required if specifying 1 (ICMP) for the protocol parameter.

*Type*: [EC2 ICMP Property Type](aws-properties-ec2-icmp.html "EC2 ICMP Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NetworkAclId`   
ID of the ACL where the entry will be created.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

 `PortRange`   
The range of port numbers for the UDP/TCP protocol.

*Required*: Conditional Required if specifying 6 (TCP) or 17 (UDP) for the protocol parameter.

*Type*: [EC2 PortRange Property Type](aws-properties-ec2-port-range.html "EC2 PortRange Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Protocol`   
The IP protocol that the rule applies to. You must specify `-1` or a protocol number (go to [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) at iana.org). You can specify `-1` for all protocols.

Note

If you specify `-1`, all ports are opened and the `PortRange` property is ignored.

*Required*: Yes

*Type*: Number

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RuleAction`   
Whether to allow or deny traffic that matches the rule; valid values are "allow" or "deny".

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RuleNumber`   
Rule number to assign to the entry (e.g., 100). This must be a positive integer from 1 to 32766.

*Required*: Yes

*Type*: Number

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myNetworkAclEntry" : {
         "Type" : "AWS::EC2::NetworkAclEntry",
         "Properties" : {
            "NetworkAclId" : { "Ref" : "myNetworkAcl" },
            "RuleNumber" : "100",
            "Protocol" : "-1",
            "RuleAction" : "allow",
            "Egress" : "true",
            "CidrBlock" : "172.16.0.0/24",
            "Icmp" : { "Code" : "-1", "Type" : "-1" },
            "PortRange" : { "From" : "53", "To" : "53" }
         }
      }
   }
}     
    
```

See Also
--------

-   [NetworkAclEntry](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAclEntry.html) in the *Amazon EC2 API Reference*

-   [Network ACLs](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html) in the *Amazon Virtual Private Cloud User Guide*.


