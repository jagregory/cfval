Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints
============================================================================

`MountPoints` is a property of the [Amazon EC2 Container Service TaskDefinition ContainerDefinitions](aws-properties-ecs-taskdefinition-containerdefinitions.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions") property that specifies the mount points for data volumes in a container.

Syntax
------

``` {.programlisting}
      {
  "ContainerPath" : String,
  "SourceVolume" : String,
  "ReadOnly" : Boolean
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `ContainerPath`   
The path on the container that indicates where you want to mount the volume.

*Required*: Yes

*Type*: String

 `SourceVolume`   
The name of the volume to mount.

*Required*: Yes

*Type*: String

 `ReadOnly`   
Indicates whether the container can write to the volume. If you specify **`true`**, the container has read-only access to the volume. If you specify **`false`**, the container can write to the volume. By default, the value is **`false`**.

*Required*: No

*Type*: Boolean


