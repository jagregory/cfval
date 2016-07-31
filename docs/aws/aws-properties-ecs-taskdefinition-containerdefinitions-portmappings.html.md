Amazon EC2 Container Service TaskDefinition ContainerDefinitions PortMappings
=============================================================================

`PortMappings` is a property of the [Amazon EC2 Container Service TaskDefinition ContainerDefinitions](aws-properties-ecs-taskdefinition-containerdefinitions.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions") property that maps a container port to a host port.

Syntax
------

``` {.programlisting}
      {
  "ContainerPort" : Integer,
  "HostPort" : Integer,
  "Protocol" : String
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `ContainerPort`   
The port number on the container that is bound to the host port.

*Required*: Yes

*Type*: Integer

 `HostPort`   
The host port number on the container instance that you want to reserve for your container. You can specify a non-reserved host port for your container port mapping, or you can omit the host port (or set it to `0`). If you specify a container port but no host port, your container port is automatically assigned a host port in the `49153` to `65535` port range.

Do not specify a host port in the `49153` to `65535` port range; these ports are reserved for automatic assignment. Other reserved ports include `22` for SSH, the Docker ports `2375` and `2376`, and the Amazon EC2 Container Service container agent port `51678`. In addition, do not specify a host port that is being used for a task; that port is reserved while the task is running.

*Required*: No

*Type*: Integer

 `Protocol`   
The protocol used for the port mapping. For valid values, see the [`protocol`](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) parameter in the *Amazon EC2 Container Service Developer Guide*. By default, AWS CloudFormation specifies `tcp`.

*Required*: No

*Type*: String


