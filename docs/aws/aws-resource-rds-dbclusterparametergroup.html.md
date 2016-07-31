AWS::RDS::DBClusterParameterGroup
=================================

The `AWS::RDS::DBClusterParameterGroup` resource creates a new Amazon Relational Database Service (Amazon RDS) database (DB) cluster parameter group. For more information about DB cluster parameter groups, see [Appendix: DB Cluster and DB Instance Parameters](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Appendix.ParameterGroups.html) in the *Amazon Relational Database Service User Guide*.

Note

Applying a parameter group to a DB cluster might require instances to reboot, resulting in a database outage while the instances reboot.

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::RDS::DBClusterParameterGroup",
  "Properties" : {
    "Description" : String,
    "Family" : String,
    "Parameters" : DBParameters,
    "Tags" : [ Resource Tag, ... ]
  }
}
    
```

Properties
----------

 `Description`   
A friendly description for this DB cluster parameter group.

*Required*: Yes

*Type:* String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Family`   
The database family of this DB cluster parameter group, such as `aurora5.6`.

*Required*: Yes

*Type:* String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Parameters`   
The parameters to set for this DB cluster parameter group. For a list of parameter keys, see [Appendix: DB Cluster and DB Instance Parameters](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Appendix.ParameterGroups.html) in the *Amazon Relational Database Service User Guide*.

Changes to dynamic parameters are applied immediately. Changes to static parameters require a reboot without failover to the DB instance that is associated with the parameter group before the change can take effect.

*Required*: Yes

*Type:* A JSON object consisting of string key-value pairs, as shown in the following example:

``` {.programlisting}
            "Parameters" : {
   "Key1" : "Value1",
   "Key2" : "Value2",
   "Key3" : "Value3"
}
          
```

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt), depending on the parameters that you update.

 `Tags`   
The tags that you want to attach to this parameter group.

*Required*: No

*Type*: A list of [resource tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: Updates are not supported.

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name..

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following snippet creates a parameter group that sets the character set database to UTF32:

``` {.programlisting}
      "RDSDBClusterParameterGroup" : {
  "Type" : "AWS::RDS::DBClusterParameterGroup",
  "Properties" : {
    "Parameters" : {
      "character_set_database" : "utf32"
    },
    "Family" : "aurora5.6",
    "Description" : "A sample parameter group"
  }
}
    
```
