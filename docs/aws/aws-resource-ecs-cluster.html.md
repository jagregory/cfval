AWS::ECS::Cluster
=================

The `AWS::ECS::Cluster` resource creates an Amazon EC2 Container Service (Amazon ECS) cluster. This resource has no properties; use the Amazon ECS container agent to connect to the cluster. For more information, see [Amazon ECS Container Agent](http://docs.aws.amazon.com/AmazonECS/latest/developerguide//ECS_agent.html) in the *Amazon EC2 Container Service Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ECS::Cluster"
}
    
```

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

In the following sample, the `Ref` function returns the name of the `MyECSCluster` cluster, such as `MyStack-MyECSCluster-NT5EUXTNTXXD`.

``` {.programlisting}
        { "Ref": "MyECSCluster" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following sample declares an Amazon ECS cluster:

``` {.programlisting}
      "MyCluster": {
  "Type": "AWS::ECS::Cluster"
}
    
```
