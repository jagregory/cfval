ElasticLoadBalancing LBCookieStickinessPolicy Type
==================================================

The LBCookieStickinessPolicy type is an embedded property of the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") type.

Syntax
------

``` {.programlisting}
      
{
   "CookieExpirationPeriod" : String,
   "PolicyName" : String
}
    
```

Properties
----------

 `CookieExpirationPeriod`   
The time period, in seconds, after which the cookie should be considered stale. If this parameter isn't specified, the sticky session will last for the duration of the browser session.

*Required*: No

*Type*: String

 `PolicyName`   
The name of the policy being created. The name must be unique within the set of policies for this load balancer.

Note

To associate this policy with a listener, include the policy name in the listener's [PolicyNames](aws-properties-ec2-elb-listener.html "ElasticLoadBalancing Listener Property Type") property.

See Also
--------

-   Sample template snippets in the Examples section of [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer").

-   [CreateLBCookieStickinessPolicy](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_CreateLBCookieStickinessPolicy.html) in the *Elastic Load Balancing API Reference*


