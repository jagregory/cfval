AWS::Route53::HostedZone
========================

The `AWS::Route53::HostedZone` resource creates a hosted zone, which can contain a collection of record sets for a domain. You cannot create a hosted zone for a top-level domain (TLD). For more information, see [POST CreateHostedZone](http://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateHostedZone.html) or [POST CreateHostedZone (Private)](http://docs.aws.amazon.com/Route53/latest/APIReference/API-create-hosted-zone-private.html) in the *Amazon Route 53 API Reference*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Route53::HostedZone",
  "Properties" : {
    "HostedZoneConfig" : { HostedZoneConfig },
    "HostedZoneTags" : [  HostedZoneTags, ... ],
    "Name" : String,
    "VPCs" : [ HostedZoneVPCs, ... ]
  }
}
    
```

Properties
----------

 `HostedZoneConfig`   
A complex type that contains an optional comment about your hosted zone.

*Required*: No

*Type*: [Amazon Route 53 HostedZoneConfig Property](aws-properties-route53-hostedzone-hostedzoneconfig.html "Amazon Route 53 HostedZoneConfig Property")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HostedZoneTags`   
An arbitrary set of tags (key–value pairs) for this hosted zone.

*Required*: No

*Type*: List of [Amazon Route 53 HostedZoneTags](aws-properties-route53-hostedzone-hostedzonetags.html "Amazon Route 53 HostedZoneTags")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Name`   
The name of the domain. For resource record types that include a domain name, specify a fully qualified domain name.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `VPCs`   
One or more VPCs that you want to associate with this hosted zone. When you specify this property, AWS CloudFormation creates a private hosted zone.

*Required*: No

*Type*: List of [Amazon Route 53 HostedZoneVPCs](aws-resource-route53-hostedzone-hostedzonevpcs.html "Amazon Route 53 HostedZoneVPCs")

If this property was specified previously and you're modifying values, updates require [no interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt). If this property wasn't specified and you add values, updates require [replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement). Also, if this property was specified and you remove all values, updates require [replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement).

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "myHostedZone" }
    
```

`Ref` returns the hosted zone ID, such as `Z23ABC4XYZL05B`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following template snippet creates a private hosted zone for the `example.com` domain.

``` {.programlisting}
      "DNS": {
  "Type": "AWS::Route53::HostedZone",
  "Properties": {
    "HostedZoneConfig": {
      "Comment": "My hosted zone for example.com"
    },
    "Name": "example.com",
    "VPCs": [{
      "VPCId": "vpc-abcd1234",
      "VPCRegion": "ap-northeast-1"
    },
    {
      "VPCId": "vpc-efgh5678",
      "VPCRegion": "us-west-2"
    }],
    "HostedZoneTags" : [{
      "Key": "SampleKey1",
      "Value": "SampleValue1"
    },
    {
      "Key": "SampleKey2",
      "Value": "SampleValue2"
    }]
  }
}
    
```
