AWS::ECS::Service
=================

The `AWS::ECS::Service` resource creates an Amazon EC2 Container Service (Amazon ECS) service that runs and maintains the desired number of tasks and associated load balancers.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ECS::Service",
  "Properties" : {
    "Cluster" : String,
    "DeploymentConfiguration" : DeploymentConfiguration,
    "DesiredCount" : Integer,
    "LoadBalancers" : [ Load Balancer Objects, ... ],
    "Role" : String,
    "TaskDefinition" : String
  }
}
    
```

Properties
----------

Note

When you use Auto Scaling or Amazon Elastic Compute Cloud (Amazon EC2) to create container instances for an Amazon ECS cluster, the Amazon ECS service resource must have a dependency on the Auto Scaling group or Amazon EC2 instances. That way the container instances are available and associated with the Amazon ECS cluster before AWS CloudFormation creates the Amazon ECS service.

 `Cluster`   
The name or Amazon Resource Name (ARN) of the cluster that you want to run your service on. If you do not specify a cluster, Amazon ECS uses the default cluster.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DeploymentConfiguration`   
Configures how many tasks run during a deployment.

*Required*: No

*Type*: [Amazon EC2 Container Service Service DeploymentConfiguration](aws-properties-ecs-service-deploymentconfiguration.html "Amazon EC2 Container Service Service DeploymentConfiguration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `DesiredCount`   
The number of simultaneous tasks, which you specify by using the `TaskDefinition` property, that you want to run on the cluster.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LoadBalancers`   
A list of load balancer objects to associate with the cluster.

*Required*: No

*Type*: List of [Amazon EC2 Container Service Service LoadBalancers](aws-properties-ecs-service-loadbalancers.html "Amazon EC2 Container Service Service LoadBalancers")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Role`   
The name or ARN of an AWS Identity and Access Management (IAM) role that allows your Amazon ECS container agent to make calls to your load balancer.

Note

In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

*Required*: Conditional. This parameter is required only if you specify the `LoadBalancers` property.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `TaskDefinition`   
The ARN of the task definition that you want to run on the cluster.

*Required*: Yes

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the ARN.

In the following sample, the `Ref` function returns the ARN of the `MyECSService` service, such as `arn:aws:ecs:us-west-2:123456789012:service/sample-webapp`.

``` {.programlisting}
        { "Ref": "MyECSService" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following sample defines an Amazon ECS service that uses a cluster and task definition that are declared elsewhere in the same template:

``` {.programlisting}
      "WebApp": {
  "Type": "AWS::ECS::Service",
  "Properties" : {
    "Cluster": { "Ref": "cluster" },
    "DesiredCount": { "Ref": "desiredcount" },
    "TaskDefinition" : { "Ref":"taskdefinition" }
  }
}
    
```

Related Resources
-----------------

For a complete sample template that shows how you can create an Amazon ECS cluster and service, see [Amazon EC2 Container Service Template Snippets](quickref-ecs.html "Amazon EC2 Container Service Template Snippets").

