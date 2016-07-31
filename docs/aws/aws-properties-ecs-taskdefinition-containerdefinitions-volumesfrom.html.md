Amazon EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom
============================================================================

`VolumesFrom` is a property of the [Amazon EC2 Container Service TaskDefinition ContainerDefinitions](aws-properties-ecs-taskdefinition-containerdefinitions.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions") property that mounts data volumes from other containers.

Syntax
------

``` {.programlisting}
      {
  "SourceContainer" : String,
  "ReadOnly" : Boolean
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `SourceContainer`   
The name of the container that has the volumes to mount.

*Required*: Yes

*Type*: String

 `ReadOnly`   
Indicates whether the container can write to the volume. If you specify **`true`**, the container has read-only access to the volume. If you specify **`false`**, the container can write to the volume. By default, the value is **`false`**.

*Required*: No

*Type*: Boolean


