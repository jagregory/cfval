AWS::EC2::PlacementGroup
========================

The `AWS::EC2::PlacementGroup` resource is a logical grouping of instances within a single Availability Zone (AZ) that enables applications to participate in a low-latency, 10 Gbps network. You create a placement group first, and then you can launch instances in the placement group.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EC2::PlacementGroup",
  "Properties" : {
    "Strategy" : String
  }
}
    
```

Properties
----------

 `Strategy`   
The placement strategy, which relates to the instance types that can be added to the placement group. For example, for the `cluster` strategy, you can cluster C4 instance types but not T2 instance types. For valid values, see [CreatePlacementGroup](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreatePlacementGroup.html) in the *Amazon EC2 API Reference*. By default, AWS CloudFormation sets the value of this property to `cluster`.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

### 

The following example creates a placement group with a `cluster` placement strategy.

``` {.programlisting}
        "PlacementGroup" : {
  "Type" : "AWS::EC2::PlacementGroup",
  "Properties" : {
    "Strategy" : "cluster"
  }
}
      
```
