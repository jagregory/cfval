Elastic Load Balancing ConnectionDrainingPolicy
===============================================

The `ConnectionDrainingPolicy` property describes how deregistered or unhealthy instances handle in-flight requests for the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") resource. Connection draining ensures that the load balancer completes serving all in-flight requests made to a registered instance when the instance is deregistered or becomes unhealthy. Without connection draining, the load balancer closes connections to deregistered or unhealthy instances, and any in-flight requests are not completed.

For more information about connection draining and default values, see [Enable or Disable Connection Draining for Your Load Balancer](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/config-conn-drain.html) in the *Elastic Load Balancing Developer Guide*.

Syntax
------

``` {.programlisting}
      
{
   "Enabled" : Boolean,
   "Timeout" : Integer
}
    
```

Properties
----------

 `Enabled`   
Whether or not connection draining is enabled for the load balancer.

*Required*: Yes

*Type*: Boolean

 `Timeout`   
The time in seconds after the load balancer closes all connections to a deregistered or unhealthy instance.

*Required*: No

*Type*: Integer


