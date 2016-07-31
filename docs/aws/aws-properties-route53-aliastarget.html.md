Route 53 AliasTarget Property
=============================

`AliasTarget` is a property of the [AWS::Route53::RecordSet](aws-properties-route53-recordset.html "AWS::Route53::RecordSet") resource.

For more information about alias resource record sets, see [Creating Alias Resource Record Sets](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/CreatingAliasRRSets.html) in the *Amazon Route 53 Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "DNSName" : String,
  "EvaluateTargetHealth" : Boolean,
  "HostedZoneId" : String
}
    
```

Properties
----------

 `DNSName`   
The DNS name of the load balancer, the domain name of the CloudFront distribution, the website endpoint of the Amazon S3 bucket, or another record set in the same hosted zone that is the target of the alias.

*Type*: String

*Required*: Yes

 `EvaluateTargetHealth`   
Whether Amazon Route 53 checks the health of the resource record sets in the alias target when responding to DNS queries. For more information about using this property, see [EvaluateTargetHealth](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-evaluate-target-health) in the *Amazon Route 53 API Reference*.

*Type*: Boolean

*Required*: No

 `HostedZoneId`   
The hosted zone ID. For load balancers, use the canonical hosted zone ID of the load balancer. For Amazon S3, use the hosted zone ID for your bucket's website endpoint. For CloudFront, use `Z2FDTNDATAQYW2`. For examples, see [Example: Creating Alias Resource Record Sets](http://docs.aws.amazon.com/Route53/latest/APIReference/CreateAliasRRSAPI.html) in the *Amazon Route 53 API Reference*.

*Type*: String

*Required*: Yes


