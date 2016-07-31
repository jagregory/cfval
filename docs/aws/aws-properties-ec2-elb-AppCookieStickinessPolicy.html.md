ElasticLoadBalancing AppCookieStickinessPolicy Type
===================================================

The AppCookieStickinessPolicy type is an embedded property of the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") type.

Syntax
------

``` {.programlisting}
      
{
   "CookieName" : String,
   "PolicyName" : String
}
    
```

Properties
----------

 `CookieName`   
Name of the application cookie used for stickiness.

*Required*: Yes

*Type*: String

 `PolicyName`   
The name of the policy being created. The name must be unique within the set of policies for this Load Balancer.

Note

To associate this policy with a listener, include the policy name in the listener's [PolicyNames](aws-properties-ec2-elb-listener.html "ElasticLoadBalancing Listener Property Type") property.

*Required*: Yes

*Type*: String

See Also
--------

-   Sample template snippets in the Examples section of [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer").

-   [CreateAppCookieStickinessPolicy](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_CreateAppCookieStickinessPolicy.html)in the *Elastic Load Balancing API Reference*


