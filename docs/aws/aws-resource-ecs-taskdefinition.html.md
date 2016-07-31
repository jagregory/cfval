AWS::ECS::TaskDefinition
========================

The `AWS::ECS::TaskDefinition` resource describes the container and volume definitions of an Amazon EC2 Container Service (Amazon ECS) task. You can specify which Docker images to use, the required resources, and other configurations related to launching the task definition through an Amazon ECS service or task.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ECS::TaskDefinition",
  "Properties" : {
    "ContainerDefinitions" : [ Container Definition, ... ],
    "Volumes" : [ Volume Definition, ... ]
  }
}
    
```

Properties
----------

 `ContainerDefinitions`   
A list of container definitions in JSON format that describe the containers that make up your task.

*Required*: Yes

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions](aws-properties-ecs-taskdefinition-containerdefinitions.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Volumes`   
A list of volume definitions in JSON format for volumes that you can use in your container definitions.

*Required*: Yes

*Type*: List of [Amazon EC2 Container Service TaskDefinition Volumes](aws-properties-ecs-taskdefinition-volumes.html "Amazon EC2 Container Service TaskDefinition Volumes")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the Amazon Resource Name (ARN).

In the following sample, the `Ref` function returns the ARN of the `MyTaskDefinition` task, such as `arn:aws:ecs:us-west-2:123456789012:task/1abf0f6d-a411-4033-b8eb-a4eed3ad252a`.

``` {.programlisting}
        { "Ref": "MyTaskDefinition" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following example defines an Amazon ECS task definition, which includes two container definitions and one volume definition:

``` {.programlisting}
      "taskdefinition": {
  "Type": "AWS::ECS::TaskDefinition",
  "Properties" : {
    "ContainerDefinitions" : [
    {
      "Name": {"Ref": "AppName"},
      "MountPoints": [
        {
          "SourceVolume": "my-vol",
          "ContainerPath": "/var/www/my-vol"
        }
      ],
      "Image":"amazon/amazon-ecs-sample",
      "Cpu": "10",
      "PortMappings":[
        {
          "ContainerPort": {"Ref":"AppContainerPort"},
          "HostPort": {"Ref":"AppHostPort"}
        }
      ],
      "EntryPoint": [
        "/usr/sbin/apache2",
        "-D",
        "FOREGROUND"
      ],
      "Memory":"500",
      "Essential": "true"
    },
    {
      "Name": "busybox",
      "Image": "busybox",
      "Cpu": "10",
      "EntryPoint": [
        "sh",
        "-c"
      ],
      "Memory": "500",
      "Command": [
        "/bin/sh -c \"while true; do /bin/date > /var/www/my-vol/date; sleep 1; done\""
      ],
      "Essential" : "false",
      "VolumesFrom": [
        {
          "SourceContainer": {"Ref":"AppName"}
        }
      ]
    }],
    "Volumes": [
    {
      "Host": {
        "SourcePath": "/var/lib/docker/vfs/dir/"
      },
      "Name": "my-vol"
    }]
  }
}
    
```

Related Resources
-----------------

For a complete sample template that shows how you can create an Amazon ECS cluster and service, see [Amazon EC2 Container Service Template Snippets](quickref-ecs.html "Amazon EC2 Container Service Template Snippets").

