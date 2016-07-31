Amazon Route 53 HostedZoneVPCs
==============================

The `HostedZoneVPCs` property is part of the [AWS::Route53::HostedZone](aws-resource-route53-hostedzone.html "AWS::Route53::HostedZone") resource that specifies the VPCs to associate with the hosted zone.

Syntax
------

``` {.programlisting}
      {
  "VPCId" : String,
  "VPCRegion" : String
}
    
```

Properties
----------

 `VPCId`   
The ID of the Amazon VPC that you want to associate with the hosted zone.

*Required*: Yes

*Type*: String

 `VPCRegion`   
The region in which the Amazon VPC was created as specified in the `VPCId` property.

*Required*: Yes

*Type*: String


