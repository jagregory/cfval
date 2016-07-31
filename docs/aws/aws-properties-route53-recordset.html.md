AWS::Route53::RecordSet
=======================

The AWS::Route53::RecordSet type can be used as a standalone resource or as an embedded property in the [AWS::Route53::RecordSetGroup](aws-properties-route53-recordsetgroup.html "AWS::Route53::RecordSetGroup") type. Note that some AWS::Route53::RecordSet properties are valid only when used within AWS::Route53::RecordSetGroup.

For more information about constraints and values for each property, see [POST CreateHostedZone](http://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateHostedZone.html) for hosted zones and [POST ChangeResourceRecordSet](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) for resource record sets.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Route53::RecordSet",
  "Properties" : {
    "AliasTarget" : AliasTarget,
    "Comment" : String,
    "Failover" : String,
    "GeoLocation" : { GeoLocation },
    "HealthCheckId" : String,
    "HostedZoneId" : String,
    "HostedZoneName" : String,
    "Name" : String,
    "Region" : String,
    "ResourceRecords" : [ String ],
    "SetIdentifier" : String,
    "TTL" : String,
    "Type" : String,
    "Weight" : Integer
  }
}
    
```

Properties
----------

 `AliasTarget`   
*Alias resource record sets only:* Information about the domain to which you are redirecting traffic.

If you specify this property, do not specify the `TTL` property. The alias uses a TTL value from the alias target record.

For more information about alias resource record sets, see [Creating Alias Resource Record Sets](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/CreatingAliasRRSets.html) in the *Amazon Route 53 Developer Guide* and [POST ChangeResourceRecordSets](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html#API_ChangeResourceRecordSets_RequestAliasSyntax) in the Amazon Route 53 API reference.

*Required*: Conditional. Required if you are creating an alias resource record set.

*Type*: [AliasTarget](aws-properties-route53-aliastarget.html "Route 53 AliasTarget Property")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Comment`   
Any comments that you want to include about the hosted zone.

Important

If the record set is part of a record set group, this property isn't valid. Don't specify this property.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Failover`   
Designates the record set as a `PRIMARY` or `SECONDARY` failover record set. When you have more than one resource performing the same function, you can configure Amazon Route 53 to check the health of your resources and use only health resources to respond to DNS queries. You cannot create nonfailover resource record sets that have the same `Name` and `Type` property values as failover resource record sets. For more information, see the [Failover](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-failover) element in the *Amazon Route 53 API Reference*.

If you specify this property, you must specify the `SetIdentifier` property.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `GeoLocation`   
Describes how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.

*Required*: No

*Type*: [Amazon Route 53 Record Set GeoLocation Property](aws-properties-route53-recordset-geolocation.html "Amazon Route 53 Record Set GeoLocation Property")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HealthCheckId`   
The health check ID that you want to apply to this record set. Amazon Route 53 returns this resource record set in response to a DNS query only while record set is healthy.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HostedZoneId`   
The ID of the hosted zone.

*Required*: Conditional. You must specify either the *`HostedZoneName`* or *`HostedZoneId`*, but you cannot specify both. If this record set is part of a record set group, do not specify this property.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `HostedZoneName`   
The name of the domain for the hosted zone where you want to add the record set.

When you create a stack using an AWS::Route53::RecordSet that specifies *`HostedZoneName`*, AWS CloudFormation attempts to find a hosted zone whose name matches the HostedZoneName. If AWS CloudFormation cannot find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name, AWS CloudFormation will not create the stack.

If you have multiple hosted zones with the same domain name, you must explicitly specify the hosted zone using *`HostedZoneId`*.

*Required*: Conditional. You must specify either the *`HostedZoneName`* or *`HostedZoneId`*, but you cannot specify both. If this record set is part of a record set group, do not specify this property.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Name`   
The name of the domain. You must specify a fully qualified domain name that ends with a period as the last label indication. If you omit the final period, AWS CloudFormation adds it.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Region`   
Latency resource record sets only: The Amazon EC2 region where the resource that is specified in this resource record set resides. The resource typically is an AWS resource, for example, Amazon EC2 instance or an Elastic Load Balancing load balancer, and is referred to by an IP address or a DNS domain name, depending on the record type.

When Amazon Route 53 receives a DNS query for a domain name and type for which you have created latency resource record sets, Amazon Route 53 selects the latency resource record set that has the lowest latency between the end user and the associated Amazon EC2 region. Amazon Route 53 then returns the value that is associated with the selected resource record set.

The following restrictions must be followed:

-   You can only specify one resource record per latency resource record set.

-   You can only create one latency resource record set for each Amazon EC2 region.

-   You are not required to create latency resource record sets for all Amazon EC2 regions. Amazon Route 53 will choose the region with the best latency from among the regions for which you create latency resource record sets.

-   You cannot create both weighted and latency resource record sets that have the same values for the Name and Type elements.

To see a list of regions by service, see [Regions and Endpoints](http://docs.aws.amazon.com/general/latest/gr/rande.html) in the *AWS General Reference*.

 `ResourceRecords`   
List of resource records to add. Each record should be in the format appropriate for the record type specified by the *`Type`* property. For information about different record types and their record formats, see [Appendix: Domain Name Format](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html) in the *Amazon Route 53 Developer Guide*.

*Required*: Conditional. If you don't specify the `AliasTarget` property, you must specify this property. If you are creating an alias resource record set, do not specify this property.

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SetIdentifier`   
A unique identifier that differentiates among multiple resource record sets that have the same combination of DNS name and type.

*Required*: Conditional. Required if you are creating a weighted, latency, failover, or geolocation resource record set.

For more information, see the [SetIdentifier](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-set-id) element in the *Amazon Route 53 Developer Guide*.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `TTL`   
The resource record cache time to live (TTL), in seconds. If you specify this property, do not specify the `AliasTarget` property. For alias target records, the alias uses a TTL value from the target.

If you specify this property, you must specify the *`ResourceRecords`* property.

*Required*: Conditional. If you don't specify the `AliasTarget` property, you must specify this property. If you are creating an alias resource record set, do not specify this property.

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Type`   
The type of records to add.

*Required*: Yes

*Type*: String

*Valid Values*: A | AAAA | CNAME | MX | NS | PTR | SOA | SPF | SRV | TXT

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Weight`   
*Weighted resource record sets only:* Among resource record sets that have the same combination of DNS name and type, a value that determines what portion of traffic for the current resource record set is routed to the associated location.

For more information about weighted resource record sets, see [Setting Up Weighted Resource Record Sets](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/WeightedResourceRecordSets.html) in the *Amazon Route 53 Developer Guide*.

*Required*: Conditional. Required if you are creating a weighted resource record set.

*Type*: Number. Weight expects integer values.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

When you specify an AWS::Route53::RecordSet type as an argument to the `Ref` function, AWS CloudFormation returns the value of the domain name of the record set.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

**Example Mapping an Amazon Route 53 A record to the public IP of an Amazon EC2 instance**

``` {.programlisting}
          
"Resources" : {
   "Ec2Instance" : {
      "Type" : "AWS::EC2::Instance",
      "Properties" : {
         "ImageId" : { "Fn::FindInMap" : [
            "RegionMap", { "Ref" : "AWS::Region" }, "AMI"
         ] }
      }
   },
   "myDNSRecord" : {
      "Type" : "AWS::Route53::RecordSet",
      "Properties" : {
         "HostedZoneName" : {
            "Fn::Join" : [ "", [
               { "Ref" : "HostedZone" }, "."
            ] ]
         },
         "Comment" : "DNS name for my instance.",  
         "Name" : {
            "Fn::Join" : [ "", [
               {"Ref" : "Ec2Instance"}, ".",
               {"Ref" : "AWS::Region"}, ".",
               {"Ref" : "HostedZone"} ,"."
            ] ]
         },
         "Type" : "A",
         "TTL" : "900",
         "ResourceRecords" : [
            { "Fn::GetAtt" : [ "Ec2Instance", "PublicIp" ] }
         ]
      }
   }
},
        
```

For additional AWS::Route53::RecordSet snippets, see [Amazon Route 53 Template Snippets](quickref-route53.html "Amazon Route 53 Template Snippets") .

