Amazon EC2 Container Service TaskDefinition Volumes
===================================================

`Volumes` is a property of the [AWS::ECS::TaskDefinition](aws-resource-ecs-taskdefinition.html "AWS::ECS::TaskDefinition") resource that specifies a list of data volumes, which your containers can then access.

Syntax
------

``` {.programlisting}
      {
  "Name" : String,
  "Host" : Host
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `Name`   
The name of the volume. To specify mount points in your container definitions, use the value of this property.

*Required*: Yes

*Type*: String

 `Host`   
Determines whether your data volume persists on the host container instance and at the location where it is stored.

*Required*: No

*Type*: [Amazon EC2 Container Service TaskDefinition Volumes Host](aws-properties-ecs-taskdefinition-volumes-host.html "Amazon EC2 Container Service TaskDefinition Volumes Host")


