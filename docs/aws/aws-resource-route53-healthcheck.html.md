AWS::Route53::HealthCheck
=========================

You can use the `AWS::Route53::HealthCheck` resource to check the health of your resources before Amazon Route 53 responds to a DNS query. For more information, see [How Health Checks Work in Simple Amazon Route 53 Configurations](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-simple-configs.html) in the *Amazon Route 53 Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Route53::HealthCheck",
  "Properties" : {
    "HealthCheckConfig" : { HealthCheckConfig },
    "HealthCheckTags" : [ HealthCheckTags, ... ]
  }
}
    
```

Properties
----------

 `HealthCheckConfig`   
An Amazon Route 53 health check.

*Required*: Yes

*Type*: [Amazon Route 53 HealthCheckConfig](aws-properties-route53-healthcheck-healthcheckconfig.html "Amazon Route 53 HealthCheckConfig")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HealthCheckTags`   
An arbitrary set of tags (key–value pairs) for this health check.

*Required*: No

*Type*: List of [Amazon Route 53 HealthCheckTags](aws-properties-route53-healthcheck-healthchecktags.html "Amazon Route 53 HealthCheckTags")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the health check ID, such as `e0a123b4-4dba-4650-935e-example`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following template snippet creates an Amazon Route 53 health check that sends request to the specified endpoint.

``` {.programlisting}
      "myHealthCheck": {
  "Type": "AWS::Route53::HealthCheck",
  "Properties": {
    "HealthCheckConfig": {
      "IPAddress": "000.000.000.000",
      "Port": "80",
      "Type": "HTTP",
      "ResourcePath": "/example/index.html",
      "FullyQualifiedDomainName": "example.com",
      "RequestInterval": "30",
      "FailureThreshold": "3"
    },
    "HealthCheckTags" : [{
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
