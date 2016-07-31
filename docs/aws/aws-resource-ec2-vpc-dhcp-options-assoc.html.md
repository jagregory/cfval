AWS::EC2::VPCDHCPOptionsAssociation
===================================

Associates a set of DHCP options (that you've previously created) with the specified VPC.

Syntax
------

``` {.programlisting}
      
{ 
   "Type" : "AWS::EC2::VPCDHCPOptionsAssociation",
   "Properties" : {
      "DhcpOptionsId" : String,
      "VpcId" : String
   }
}     
    
```

Properties
----------

 `DhcpOptionsId`   
The ID of the DHCP options you want to associate with the VPC. Specify `default` if you want the VPC to use no DHCP options.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VpcId`   
The ID of the VPC to associate with this DHCP options set.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following snippet uses the `Ref` intrinsic function to associate the `myDHCPOptions` DHCP options with the `myVPC` VPC. The VPC and DHCP options can be declared in the same template or added as input parameters. For more information about the VPC or the DHCP options resources, see [AWS::EC2::VPC](aws-resource-ec2-vpc.html "AWS::EC2::VPC") or [AWS::EC2::DHCPOptions](aws-resource-ec2-dhcp-options.html "AWS::EC2::DHCPOptions").

``` {.programlisting}
      "myVPCDHCPOptionsAssociation" : {
  "Type" : "AWS::EC2::VPCDHCPOptionsAssociation",
  "Properties" : {
    "VpcId" : {"Ref" : "myVPC"},
    "DhcpOptionsId" : {"Ref" : "myDHCPOptions"}
  }
}
    
```

See Also
--------

-   [AssociateDhcpOptions](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateDhcpOptions.html) in the *Amazon EC2 API Reference*.


