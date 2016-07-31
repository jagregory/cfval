Condition Functions
===================

You can use intrinsic functions, such as `Fn::If`, `Fn::Equals`, and `Fn::Not`, to conditionally create stack resources. These conditions are evaluated based on input parameters that you declare when you create or update a stack. After you define all your conditions, you can associate them with resources or resource properties in the Resources and Outputs sections of a template.

You define all conditions in the Conditions section of a template except for `Fn::If` conditions. You can use the `Fn::If` condition in the metadata attribute, update policy attribute, and property values in the Resources section and Outputs sections of a template.

You might use conditions when you want to reuse a template that can create resources in different contexts, such as a test environment versus a production environment. In your template, you can add an `EnvironmentType` input parameter, which accepts either **`prod`** or **`test`** as inputs. For the production environment, you might include Amazon EC2 instances with certain capabilities; however, for the test environment, you want to use less capabilities to save costs. With conditions, you can define which resources are created and how they're configured for each environment type.

For more information about the Conditions section, see [Conditions](conditions-section-structure.html "Conditions").

Note

You can only reference other conditions and values from the Parameters and Mappings sections of a template. For example, you can reference a value from an input parameter, but you cannot reference the logical ID of a resource in a condition.

Associating a Condition

To conditionally create resources, resource properties, or outputs, you must associate a condition with them. Add the `Condition:` key and the logical ID of the condition as an attribute to associate a condition, as shown in the following snippet. AWS CloudFormation creates the `NewVolume` resource only when the `CreateProdResources` condition evaluates to true.

``` {.programlisting}
    "NewVolume" : {
  "Type" : "AWS::EC2::Volume",
  "Condition" : "CreateProdResources",
  "Properties" : {
     "Size" : "100",
     "AvailabilityZone" : { "Fn::GetAtt" : [ "EC2Instance", "AvailabilityZone" ]}
}
  
```

For the `Fn::If` function, you only need to specify the condition name. The following snippet shows how to use `Fn::If` to conditionally specify a resource property. If the `CreateLargeSize` condition is true, AWS CloudFormation sets the volume size to `100`. If the condition is false, AWS CloudFormation sets the volume size to `10`.

``` {.programlisting}
    "NewVolume" : {
  "Type" : "AWS::EC2::Volume",
  "Properties" : {
    "Size" : {
      "Fn::If" : [
        "CreateLargeSize",
        "100",
        "10"
      ]},
    "AvailabilityZone" : { "Fn::GetAtt" : [ "Ec2Instance", "AvailabilityZone" ]}
  },
  "DeletionPolicy" : "Snapshot"
}
  
```

You can also use conditions inside other conditions. The following snippet is from the `Conditions` section of a template. The `MyAndCondition` condition includes the `SomeOtherCondition` condition:

``` {.programlisting}
    "MyAndCondition": {
   "Fn::And": [
      {"Fn::Equals": ["sg-mysggroup", {"Ref": "ASecurityGroup"}]},
      {"Condition": "SomeOtherCondition"}
   ]
}
  
```

**Topics**

-   [Fn::And](intrinsic-function-reference-conditions.html#d0e121706)
-   [Fn::Equals](intrinsic-function-reference-conditions.html#d0e121788)
-   [Fn::If](intrinsic-function-reference-conditions.html#d0e121863)
-   [Fn::Not](intrinsic-function-reference-conditions.html#d0e122042)
-   [Fn::Or](intrinsic-function-reference-conditions.html#d0e122130)
-   [Supported Functions](intrinsic-function-reference-conditions.html#d0e122214)
-   [Sample Templates](conditions-sample-templates.html)

Fn::And
-------

Returns `true` if all the specified conditions evaluate to true, or returns `false` if any one of the conditions evaluates to false. `Fn::And` acts as an AND operator. The minimum number of conditions that you can include is 2, and the maximum is 10.

### Declaration

``` {.programlisting}
        "Fn::And": [{condition}, {...}]
      
```

### Parameters

 `condition`   
A condition that evaluates to `true` or `false`.

### Example

The following `MyAndCondition` evaluates to true if the referenced security group name is equal to `sg-mysggroup` and if `SomeOtherCondition` evaluates to true:

``` {.programlisting}
        "MyAndCondition": {
   "Fn::And": [
      {"Fn::Equals": ["sg-mysggroup", {"Ref": "ASecurityGroup"}]},
      {"Condition": "SomeOtherCondition"}
   ]
}
      
```

Fn::Equals
----------

Compares if two values are equal. Returns `true` if the two values are equal or `false` if they aren't.

### Declaration

``` {.programlisting}
        "Fn::Equals" : ["value_1", "value_2"]
      
```

### Parameters

 `value`   
A value of any type that you want to compare.

### Example

The following `UseProdCondition` condition evaluates to true if the value for the `EnvironmentType` parameter is equal to `prod`:

``` {.programlisting}
        "UseProdCondition" : {
   "Fn::Equals": [
      {"Ref": "EnvironmentType"},
      "prod"
   ]
}
      
```

Fn::If
------

Returns one value if the specified condition evaluates to `true` and another value if the specified condition evaluates to `false`. Currently, AWS CloudFormation supports the `Fn::If` intrinsic function in the metadata attribute, update policy attribute, and property values in the Resources section and Outputs sections of a template. You can use the `AWS::NoValue` pseudo parameter as a return value to remove the corresponding property.

### Declaration

``` {.programlisting}
        "Fn::If": [condition_name, value_if_true, value_if_false]
      
```

### Parameters

 `condition_name`   
A reference to a condition in the Conditions section. Use the condition's name to reference it.

 `value_if_true`   
A value to be returned if the specified condition evaluates to `true`.

 `value_if_false`   
A value to be returned if the specified condition evaluates to `false`.

### Examples

The following snippet uses an `Fn::If` function in the `SecurityGroups` property for an Amazon EC2 resource. If the `CreateNewSecurityGroup` condition evaluates to true, AWS CloudFormation uses the referenced value of `NewSecurityGroup` to specify the `SecurityGroups` property; otherwise, AWS CloudFormation uses the referenced value of `ExistingSecurityGroup`.

``` {.programlisting}
        "SecurityGroups" : [{
  "Fn::If" : [
    "CreateNewSecurityGroup",
    {"Ref" : "NewSecurityGroup"},
    {"Ref" : "ExistingSecurityGroup"}
  ]
}]
      
```

In the Output section of a template, you can use the `Fn::If` function to conditionally output information. In the following snippet, if the `CreateNewSecurityGroup` condition evaluates to true, AWS CloudFormation outputs the security group ID of the `NewSecurityGroup` resource. If the condition is false, AWS CloudFormation outputs the security group ID of the `ExistingSecurityGroup` resource.

``` {.programlisting}
        "Outputs" : {
  "SecurityGroupId" : {
    "Description" : "Group ID of the security group used.",
    "Value" : {
      "Fn::If" : [
        "CreateNewSecurityGroup",
        {"Ref" : "NewSecurityGroup"},
        {"Ref" : "ExistingSecurityGroup"}
      ]
    }
  }
}
      
```

The following snippet uses the `AWS::NoValue` pseudo parameter in an `Fn::If` function. The condition uses a snapshot for an Amazon RDS DB instance only if a snapshot ID is provided. If the `UseDBSnapshot` condition evaluates to true, AWS CloudFormation uses the `DBSnapshotName` parameter value for the `DBSnapshotIdentifier` property. If the condition evaluates to false, AWS CloudFormation removes the `DBSnapshotIdentifier` property.

``` {.programlisting}
        "MyDB" : {
  "Type" : "AWS::RDS::DBInstance",
  "Properties" : {
    "AllocatedStorage" : "5",
    "DBInstanceClass" : "db.m1.small",
    "Engine" : "MySQL",
    "EngineVersion" : "5.5",
    "MasterUsername" : { "Ref" : "DBUser" },
    "MasterUserPassword" : { "Ref" : "DBPassword" },
    "DBParameterGroupName" : { "Ref" : "MyRDSParamGroup" },
    "DBSnapshotIdentifier" : {
      "Fn::If" : [
        "UseDBSnapshot",
        {"Ref" : "DBSnapshotName"},
        {"Ref" : "AWS::NoValue"}
      ]
    }
  }
}
      
```

The following snippet provides an auto scaling update policy only if the `RollingUpdates` condition evaluates to true. If the condition evaluates to false, AWS CloudFormation removes the `AutoScalingRollingUpdate` update policy.

``` {.programlisting}
        "UpdatePolicy": {
  "AutoScalingRollingUpdate": {
    "Fn::If": [
      "RollingUpdates",
      {
        "MaxBatchSize": "2",
        "MinInstancesInService": "2",
        "PauseTime": "PT0M30S"
      },
      {
        "Ref" : "AWS::NoValue"
      }  
    ]
  }
}
      
```

To view additional samples, see [Sample Templates](conditions-sample-templates.html "Sample Templates").

Fn::Not
-------

Returns `true` for a condition that evaluates to `false` or returns `false` for a condition that evaluates to `true`. `Fn::Not` acts as a NOT operator.

### Declaration

``` {.programlisting}
        "Fn::Not": [{condition}]
      
```

### Parameters

 `condition`   
A condition such as `Fn::Equals` that evaluates to `true` or `false`.

### Example

The following `EnvCondition` condition evaluates to true if the value for the `EnvironmentType` parameter is not equal to `prod`:

``` {.programlisting}
        "MyNotCondition" : {
   "Fn::Not" : [{
      "Fn::Equals" : [
         {"Ref" : "EnvironmentType"},
         "prod"
      ]
   }]
}
      
```

Fn::Or
------

Returns `true` if any one of the specified conditions evaluate to true, or returns `false` if all of the conditions evaluates to false. `Fn::Or` acts as an OR operator. The minimum number of conditions that you can include is 2, and the maximum is 10.

### Declaration

``` {.programlisting}
        "Fn::Or": [{condition}, {...}]
      
```

### Parameters

 `condition`   
A condition that evaluates to `true` or `false`.

### Example

The following `MyOrCondition` evaluates to true if the referenced security group name is equal to `sg-mysggroup` or if `SomeOtherCondition` evaluates to true:

``` {.programlisting}
        "MyOrCondition" : {
   "Fn::Or" : [
      {"Fn::Equals" : ["sg-mysggroup", {"Ref" : "ASecurityGroup"}]},
      {"Condition" : "SomeOtherCondition"}
   ]
}
      
```

Supported Functions
-------------------

You can use the following functions in the `Fn::If` condition:

-   `Fn::Base64`

-   `Fn::FindInMap`

-   `Fn::GetAtt`

-   `Fn::GetAZs`

-   `Fn::If`

-   `Fn::Join`

-   `Fn::Select`

-   `Ref`

You can use the following functions in all other condition functions, such as `Fn::Equals` and `Fn::Or`:

-   `Fn::FindInMap`

-   `Ref`

-   Other condition functions


