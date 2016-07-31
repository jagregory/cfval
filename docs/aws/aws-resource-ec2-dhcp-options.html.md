AWS::EC2::DHCPOptions
=====================

Creates a set of DHCP options for your VPC.

For more information, see [CreateDhcpOptions](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateDhcpOptions.html) in the *Amazon EC2 API Reference*.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::DHCPOptions",
   "Properties" : {
      "DomainName" : String,
      "DomainNameServers" : [ String, ... ],
      "NetbiosNameServers" : [ String, ... ],
      "NetbiosNodeType" : Number,
      "NtpServers" : [ String, ... ],
      "Tags" : [ Resource Tag, ... ]
   }
}     
    
```

Properties
----------

 `DomainName`   
A domain name of your choice.

*Required*: Conditional; see [note](aws-resource-ec2-dhcp-options.html#dhcp-options-conditional-note "Conditional Properties").

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `"example.com"`

 `DomainNameServers`   
The IP (IPv4) address of a domain name server. You can specify up to four addresses.

*Required*: Conditional; see [note](aws-resource-ec2-dhcp-options.html#dhcp-options-conditional-note "Conditional Properties").

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `"DomainNameServers" : [ "10.0.0.1", "10.0.0.2"                   ]`

*Example*: To preserve the order of IP addresses, specify a comma delimited list as a single string: `"DomainNameServers" : [ "10.0.0.1, 10.0.0.2" ]`

 `NetbiosNameServers`   
The IP address (IPv4) of a NetBIOS name server. You can specify up to four addresses.

*Required*: Conditional; see [note](aws-resource-ec2-dhcp-options.html#dhcp-options-conditional-note "Conditional Properties").

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `"NetbiosNameServers" : [ "10.0.0.1", "10.0.0.2" ]`

*Example*: To preserve the order of IP addresses, specify a comma delimited list as a single string: `"NetbiosNameServers" : [ "10.0.0.1, 10.0.0.2" ]`

 `NetbiosNodeType`   
An integer value indicating the NetBIOS node type:

-   **1**: Broadcast ("B")

-   **2**: Point-to-point ("P")

-   **4**: Mixed mode ("M")

-   **8**: Hybrid ("H")

For more information about these values and about NetBIOS node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt), [RFC 1001](http://tools.ietf.org/rfc/rfc1001.txt), and [RFC 1002](http://tools.ietf.org/rfc/rfc1002.txt). We recommend that you use only the value `2` at this time (broadcast and multicast are not currently supported).

*Required*: Required if `NetBiosNameServers` is specified; optional otherwise.

*Type*: List of numbers

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `"NetbiosNodeType" : 2`

 `NtpServers`   
The IP address (IPv4) of a Network Time Protocol (NTP) server. You can specify up to four addresses.

*Required*: Conditional; see [note](aws-resource-ec2-dhcp-options.html#dhcp-options-conditional-note "Conditional Properties").

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `"NtpServers" : [ "10.0.0.1" ]`

*Example*: To preserve the order of IP addresses, specify a comma delimited list as a single string: `"NtpServers" : [ "10.0.0.1, 10.0.0.2" ]`

 `Tags`   
An arbitrary set of tags (key–value pairs) for this resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Conditional Properties
----------------------

*At least one* of the following properties must be specified:

-   [DomainNameServers](aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-domainnameservers)

-   [NetbiosNameServers](aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-netbiosnameservers)

-   [NtpServers](aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-ntpservers)

After this condition has been fulfilled, the rest of these properties are optional.

If you specify NetbiosNameServers, then NetbiosNodeType is required.

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
      "myDhcpOptions" : {
         "Type" : "AWS::EC2::DHCPOptions",
         "Properties" : {
            "DomainName" : "example.com",
            "DomainNameServers" : [ "AmazonProvidedDNS" ],
            "NtpServers" : [ "10.2.5.1" ],
            "NetbiosNameServers" : [ "10.2.5.1" ],
            "NetbiosNodeType" : 2,
            "Tags" : [ { "Key" : "foo", "Value" : "bar" } ]
         }
      }
   }
}
    
```

See Also
--------

-   [CreateDhcpOptions](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateDhcpOptions.html) in the *Amazon EC2 API Reference*

-   [Using Tags](http://docs.aws.amazon.com/AWSEC2/latest/DeveloperGuide/Using_Tags.html) in the *Amazon Elastic Compute Cloud User Guide*.

-   [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt) - *DHCP Options and BOOTP Vendor Extensions*, Network Working Group, 1997

-   [RFC 1001](http://tools.ietf.org/rfc/rfc1001.txt) - *Protocol Standard for a NetBIOS Service on a TCP/UDP Transport: Concepts and Methods*, Network Working Group, 1987

-   [RFC 1002](http://tools.ietf.org/rfc/rfc1002.txt) - *Protocol Standard for a NetBIOS Service on a TCP/UDP Transport: Detailed Specifications*, Network Working Group, 1987


