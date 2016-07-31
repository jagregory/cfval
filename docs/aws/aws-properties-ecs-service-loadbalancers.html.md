Amazon EC2 Container Service Service LoadBalancers
==================================================

`LoadBalancers` is a property of the [AWS::ECS::Service](aws-resource-ecs-service.html "AWS::ECS::Service") resource that specifies the load balancer to associate with an Amazon EC2 Container Service (Amazon ECS) service.

Syntax
------

``` {.programlisting}
      {
  "ContainerName" : String,
  "ContainerPort" : Integer,
  "LoadBalancerName" : String
}
    
```

Properties
----------

 `ContainerName`   
The name of a container to use with the load balancer.

*Required*: No

*Type*: String

 `ContainerPort`   
The port number on the container to direct load balancer traffic to. Your container instances must allow ingress traffic on this port.

*Required*: Yes

*Type*: Integer

 `LoadBalancerName`   
The name of the load balancer to associated with the Amazon ECS service.

*Required*: No

*Type*: String


