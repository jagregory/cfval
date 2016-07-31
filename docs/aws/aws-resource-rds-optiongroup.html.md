AWS::RDS::OptionGroup
=====================

Use the `AWS::RDS::OptionGroup` resource to create an option group that can make managing data and databases easier. For more information about option groups, see [Working with Option Groups](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html) in the *Amazon Relational Database Service User Guide*.

Syntax
------

``` {.programlisting}
      {
   "Type": "AWS::RDS::OptionGroup",
   "Properties" : {
      "EngineName" : String,
      "MajorEngineVersion" : String,
      "OptionGroupDescription" : String,
      "OptionConfigurations" : [ OptionConfigurations, ... ],
      "Tags" : [ Resource Tag, ... ]
   }
}
    
```

Properties
----------

 `EngineName`   
The name of the database engine that this option group is associated with.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `MajorEngineVersion`   
The major version number of the database engine that this option group is associated with.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `OptionGroupDescription`   
A description of the option group.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `OptionConfigurations`   
The configurations for this option group.

*Required*: Yes

*Type*: [Amazon RDS OptionGroup OptionConfigurations](aws-properties-rds-optiongroup-optionconfigurations.html "Amazon RDS OptionGroup OptionConfigurations")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this option group.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
        { "Ref": "myOptionGroup" }
      
```

For the `myOptionGroup` resource, `Ref` returns the name of the option group.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following snippet creates an option group with two option configurations (`OEM` and `APEX`):

``` {.programlisting}
      "OracleOptionGroup": {
  "Type": "AWS::RDS::OptionGroup",
  "Properties": {
    "EngineName": "oracle-ee",
    "MajorEngineVersion": "12.1",
    "OptionGroupDescription": "A test option group",
    "OptionConfigurations":[
      {
        "OptionName": "OEM",
        "DBSecurityGroupMemberships": ["default"],
        "Port": "5500"
      },
      {
        "OptionName": "APEX"
      }
    ]
  }
}
    
```

The following snippet creates an option group that specifies two option settings for the `MEMCACHED` option:

``` {.programlisting}
      "SQLOptionGroup": {
  "Type": "AWS::RDS::OptionGroup",
  "Properties": {
    "EngineName": "mysql",
    "MajorEngineVersion": "5.6",
    "OptionGroupDescription": "A test option group",
    "OptionConfigurations":[
      {
        "OptionName": "MEMCACHED",
        "VpcSecurityGroupMemberships": ["sg-a1238db7"],
        "Port": "1234",
        "OptionSettings": [
          {"Name": "CHUNK_SIZE", "Value": "32"},
          {"Name": "BINDING_PROTOCOL", "Value": "ascii"}
        ]
      }
    ]
  }
}
    
```
