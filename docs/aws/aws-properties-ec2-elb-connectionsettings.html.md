Elastic Load Balancing ConnectionSettings
=========================================

`ConnectionSettings` is a property of the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") resource that describes how long the front-end and back-end connections of your load balancer can remain idle. For more information, see [Configure Idle Connection Timeout](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/config-idle-timeout.html) in the *Elastic Load Balancing Developer Guide*.

Syntax
------

``` {.programlisting}
      
{
   "IdleTimeout" : Integer
}
    
```

Properties
----------

 `IdleTimeout`   
The time (in seconds) that a connection to the load balancer can remain idle, which means no data is sent over the connection. After the specified time, the load balancer closes the connection.

*Required*: Yes

*Type*: Integer


