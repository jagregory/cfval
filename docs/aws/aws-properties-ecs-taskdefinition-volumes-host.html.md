Amazon EC2 Container Service TaskDefinition Volumes Host
========================================================

`Host` is a property of the [Amazon EC2 Container Service TaskDefinition Volumes](aws-properties-ecs-taskdefinition-volumes.html "Amazon EC2 Container Service TaskDefinition Volumes") property that specifies the data volume path on the host container instance.

Syntax
------

``` {.programlisting}
      {
  "SourcePath" : String
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `SourcePath`   
The data volume path on the host container instance.

If you don't specify this parameter, the Docker daemon assigns a path for you, but the data volume might not persist after the associated container stops running. If you do specify a path, the data volume persists at that location on the host container instance until you manually delete it.

*Required*: No

*Type*: String


