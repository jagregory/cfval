Amazon S3 Website Configuration Routing Rules Property
======================================================

The `RoutingRules` property is an embedded property of the [Amazon S3 Website Configuration Property](aws-properties-s3-websiteconfiguration.html "Amazon S3 Website Configuration Property") property. This property describes the redirect behavior and when a redirect is applied.

Syntax
------

``` {.programlisting}
      "RoutingRules" : {
   "RedirectRule" : Redirect rule,
   "RoutingRuleCondition" : Routing rule condition
}
    
```

Properties
----------

 `                                RedirectRule                            `   
Redirect requests to another host, to another page, or with another protocol.

*Required*: Yes

*Type*: [Amazon S3 Website Configuration Routing Rules Redirect Rule Property](aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html "Amazon S3 Website Configuration Routing Rules Redirect Rule Property")

 `                                RoutingRuleCondition                            `   
Rules that define when a redirect is applied.

*Required*: No

*Type*: [Amazon S3 Website Configuration Routing Rules Routing Rule Condition Property](aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html "Amazon S3 Website Configuration Routing Rules Routing Rule Condition Property")


