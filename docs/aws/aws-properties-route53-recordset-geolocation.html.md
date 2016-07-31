Amazon Route 53 Record Set GeoLocation Property
===============================================

The `GeoLocation` property is part of the [AWS::Route53::RecordSet](aws-properties-route53-recordset.html "AWS::Route53::RecordSet") resource that describes how Amazon Route 53 responds to DNS queries based on the geographic location of the query.

Syntax
------

``` {.programlisting}
      {
  "ContinentCode" : String,
  "CountryCode" : String,
  "SubdivisionCode" : String
}
    
```

Properties
----------

 `ContinentCode`   
All DNS queries from the continent that you specified are routed to this resource record set. If you specify this property, omit the `CountryCode` and `SubdivisionCode` properties.

For valid values, see the [ContinentCode](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-continent-code) element in the *Amazon Route 53 API Reference*.

*Type*: String

*Required*: Conditional. You must specify this or the `CountryCode` property.

 `CountryCode`   
All DNS queries from the country that you specified are routed to this resource record set. If you specify this property, omit the `ContinentCode` property.

For valid values, see the [CountryCode](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-country-code) element in the *Amazon Route 53 API Reference*.

*Type*: String

*Required*: Conditional. You must specify this or the `ContinentCode` property.

 `SubdivisionCode`   
If you specified `US` for the country code, you can specify a state in the United States. All DNS queries from the state that you specified are routed to this resource record set. If you specify this property, you must specify `US` for the `CountryCode` and omit the `ContinentCode` property.

For valid values, see the [SubdivisionCode](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-subdivision-code) element in the *Amazon Route 53 API Reference*.

*Type*: String

*Required*: No


