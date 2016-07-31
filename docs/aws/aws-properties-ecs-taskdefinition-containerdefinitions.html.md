Amazon EC2 Container Service TaskDefinition ContainerDefinitions
================================================================

`ContainerDefinitions` is a property of the [AWS::ECS::TaskDefinition](aws-resource-ecs-taskdefinition.html "AWS::ECS::TaskDefinition") resource that describes the configuration of an Amazon EC2 Container Service (Amazon ECS) container. The container definitions are passed to the Docker daemon.

Syntax
------

``` {.programlisting}
      {
  "Command" : [ String, ... ],
  "Cpu" : Integer,
  "DisableNetworking" : Boolean,
  "DnsSearchDomains" : [ String, ... ],
  "DnsServers" : [ String, ... ],
  "DockerLabels" : { String:String, ... },
  "DockerSecurityOptions" : [ String, ... ],
  "EntryPoint" : [ String, ... ],
  "Environment" : [ Environment Variable, ... ],
  "Essential" : Boolean,
  "ExtraHosts" : [ Host Entry, ... ],
  "Hostname" : String,
  "Image" : String,
  "Links" : [ String, ... ],
  "LogConfiguration" : Log Configuration,
  "Memory" : Integer,
  "MountPoints" : [ Mount Point, ... ],
  "Name" : String,
  "PortMappings" : [ Port Map, ... ],
  "Privileged" : Boolean,
  "ReadonlyRootFilesystem" : Boolean,
  "Ulimits" : [ Ulimit, ... ],
  "User" : String,
  "VolumesFrom" : [ Volume From, ... ],
  "WorkingDirectory" : String,
}
    
```

Properties
----------

For more information about each property, see [Task Definition Parameters](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html) in the *Amazon EC2 Container Service Developer Guide*.

 `Command`   
The `CMD` value to pass to the container. For more information about the Docker `CMD` parameter, see <https://docs.docker.com/reference/builder/#cmd>.

*Required*: No

*Type*: List of strings

 `Cpu`   
The minimum number of CPU units to reserve for the container. Containers share unallocated CPU units with other containers on the instance by using the same ratio as their allocated CPU units. For more information, see the `cpu` content for the [ContainerDefinition](http://docs.aws.amazon.com/AmazonECS/latest/APIReference//API_ContainerDefinition.html) data type in the *Amazon EC2 Container Service API Reference*.

*Required*: No

*Type*: Integer

 `DisableNetworking`   
Indicates whether networking is disabled within the container.

*Required*: No

*Type*: Boolean

 `DnsSearchDomains`   
A list of DNS search domains that are provided to the container. The domain names that the DNS logic looks up when a process attempts to access a bare unqualified hostname.

*Required*: No

*Type*: List of strings

 `DnsServers`   
A list of DNS servers that Amazon ECS provides to the container.

*Required*: No

*Type*: List of strings

 `DockerLabels`   
A key-value map of labels for the container.

*Required*: No

*Type*: Key-value pairs, with the name of the label as the key and the label value as the value.

 `DockerSecurityOptions`   
A list of custom labels for SELinux and AppArmor multi-level security systems. For more information, see the `dockerSecurityOptions` content for the [ContainerDefinition](http://docs.aws.amazon.com/AmazonECS/latest/APIReference//API_ContainerDefinition.html) data type in the *Amazon EC2 Container Service API Reference*.

*Required*: No

*Type*: List of strings

 `EntryPoint`   
The `ENTRYPOINT` value to pass to the container. For more information about the Docker `ENTRYPOINT` parameter, see <https://docs.docker.com/reference/builder/#entrypoint>.

*Required*: No

*Type*: List of strings

 `Environment`   
The environment variables to pass to the container.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment](aws-properties-ecs-taskdefinition-containerdefinitions-environment.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment")

 `Essential`   
Indicates whether the task stops if this container fails. If you specify **`true`** and the container fails, all other containers in the task stop. If you specify **`false`** and the container fails, none of the other containers in the task is affected. This value is **`true`** by default.

You must have at least one essential container in a task.

*Required*: No

*Type*: Boolean

 `ExtraHosts`   
A list of hostnames and IP address mappings to append to the `/etc/hosts` file on the container.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions HostEntry](aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions HostEntry")

 `Hostname`   
The name that Docker will use for the container's hostname.

*Required*: No

*Type*: String

 `Image`   
The image to use for a container, which is passed directly to the Docker daemon. You can use images in the Docker Hub registry or specify other repositories (*`repository-url`*/*`image`*:*`tag`*).

*Required*: Yes

*Type*: String

 `Links`   
The name of another container to connect to. With links, containers can communicate with each other without using port mappings.

*Required*: No

*Type*: List of strings

 `LogConfiguration`   
Configures a custom log driver for the container. For more information, see the `logConfiguration` content for the [ContainerDefinition](http://docs.aws.amazon.com/AmazonECS/latest/APIReference//API_ContainerDefinition.html) data type in the *Amazon EC2 Container Service API Reference*.

*Required*: No

*Type*: [Amazon EC2 Container Service TaskDefinition ContainerDefinitions LogConfiguration](aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions LogConfiguration")

 `Memory`   
The number of MiB of memory to reserve for the container. If your container attempts to exceed the allocated memory, the container is terminated.

*Required*: Yes

*Type*: Integer

 `MountPoints`   
The mount points for data volumes in the container.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints](aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints")

 `Name`   
A name for the container.

*Required*: Yes

*Type*: String

 `PortMappings`   
A mapping of the container port to a host port. Port mappings enable containers to access ports on the host container instance to send or receive traffic.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions PortMappings](aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions PortMappings")

 `Privileged`   
Indicates whether the container is given full access to the host container instance.

*Required*: No

*Type*: Boolean

 `ReadonlyRootFilesystem`   
Indicates whether the container's root file system is mounted as read only.

*Required*: No

*Type*: Boolean

 `Ulimits`   
A list of ulimits to set in the container. The ulimits set constraints on how much resources a container can consume so that it doesn't deplete all available resources on the host.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions Ulimit](aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions Ulimit")

 `User`   
The user name to use inside the container.

*Required*: No

*Type*: String

 `VolumesFrom`   
The data volumes to mount from another container.

*Required*: No

*Type*: List of [Amazon EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom](aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html "Amazon EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom")

 `WorkingDirectory`   
The working directory in the container in which to run commands.

*Required*: No

*Type*: String


