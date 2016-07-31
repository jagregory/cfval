Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment
============================================================================

`Environment` is a property of the [Amazon EC2 Container Service TaskDefinition ContainerDefinitions](aws-properties-ecs-taskdefinition-containerdefinitions.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions") property that specifies environment variables for a container.

Syntax
------

``` {.programlisting}
      {
  "Name" : String,
  "Value" : String 
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `Name`   
The name of the environment variable.

*Required*: Yes

*Type*: String

 `Value`   
The value of the environment variable.

*Required*: Yes

*Type*: String


