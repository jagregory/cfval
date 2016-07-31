Amazon Route 53 HealthCheckConfig
=================================

The `HealthCheckConfig` property is part of the [AWS::Route53::HealthCheck](aws-resource-route53-healthcheck.html "AWS::Route53::HealthCheck") resource that describes a health check that Amazon Route 53 uses before responding to a DNS query.

Syntax
------

``` {.programlisting}
      {
  "FailureThreshold" : Integer,
  "FullyQualifiedDomainName" : String,
  "IPAddress" : String,
  "Port" : Integer,
  "RequestInterval" : Integer,
  "ResourcePath" : String,
  "SearchString" : String,
  "Type" : String
}
    
```

Properties
----------

 `FailureThreshold`   
The number of consecutive health checks that an endpoint must pass or fail for Amazon Route 53 to change the current status of the endpoint from unhealthy to healthy or healthy to unhealthy. For more information, see [How Amazon Route 53 Determines Whether an Endpoint Is Healthy](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html) in the *Amazon Route 53 Developer Guide*.

*Required*: No

*Type*: Integer

 `FullyQualifiedDomainName`   
If you specified the `IPAddress` property, the value that you want Amazon Route 53 to pass in the host header in all health checks except for TCP health checks. If you don't specify an IP address, the domain that Amazon Route 53 sends a DNS request to. Amazon Route 53 uses the IP address that the DNS returns to check the health of the endpoint.

*Required*: Conditional

*Type*: String

 `IPAddress`   
The IPv4 IP address of the endpoint on which you want Amazon Route 53 to perform health checks. If you don't specify an IP address, Amazon Route 53 sends a DNS request to resolve the domain name that you specify in the `FullyQualifiedDomainName` property.

*Required*: No

*Type*: String

 `Port`   
The port on the endpoint on which you want Amazon Route 53 to perform health checks.

*Required*: Conditional. Required when you specify `TCP` for the `Type` property.

*Type*: Integer

 `RequestInterval`   
The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request. Each Amazon Route 53 health checker makes requests at this interval. For valid values, see the [RequestInterval element](http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetHealthCheck.html) in the *Amazon Route 53 API Reference*.

*Required*: No

*Type*: Integer

 `ResourcePath`   
The path that you want Amazon Route 53 to request when performing health checks. The path can be any value for which your endpoint returns an HTTP status code of `2xx` or `3xx` when the endpoint is healthy, such as `/docs/route53-health-check.html`.

*Required*: No

*Type*: String

 `SearchString`   
If the value of the `Type` property is `HTTP_STR_MATCH` or `HTTPS_STR_MATCH`, the string that you want Amazon Route 53 to search for in the response body from the specified resource. If the string appears in the response body, Amazon Route 53 considers the resource healthy.

*Required*: No

*Type*: String

 `Type`   
The type of health check that you want to create, which indicates how Amazon Route 53 determines whether an endpoint is healthy. You can specify `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, or `TCP`. For information about the different types, see the [`Type`](http://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateHealthCheck.html) element in the *Amazon Route 53 API Reference*.

*Required*: Yes

*Type*: String


