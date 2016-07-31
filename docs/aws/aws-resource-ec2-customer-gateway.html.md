AWS::EC2::CustomerGateway
=========================

Provides information to AWS about your VPN customer gateway device.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::CustomerGateway",
   "Properties" : {
      "BgpAsn" : Number,
      "IpAddress" : String,
      "Tags" :  [ Resource Tag, ... ],
      "Type" : String
   }
}     
    
```

Properties
----------

 `BgpAsn`   
The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).

*Required*: Yes

*Type*: Number BgpAsn is always an integer value.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `IpAddress`   
The internet-routable IP address for the customer gateway's outside interface. The address must be static.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
The tags that you want to attach to the resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `Type`   
The type of VPN connection that this customer gateway supports.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

*Example*: `ipsec.1`

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "MyResource" }
    
```

For the resource with the logical ID "MyResource", `Ref` will return the AWS resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myCustomerGateway" : {
         "Type" : "AWS::EC2::CustomerGateway",
         "Properties" : {
            "Type" : "ipsec.1",
            "BgpAsn" : "64000",
            "IpAddress" : "1.1.1.1"
         }
      }
   }
}     
    
```

See Also
--------

-   [CreateCustomerGateway](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateCustomerGateway.html) in the *Amazon EC2 API Reference*.


