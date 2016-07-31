AWS::RDS::DBParameterGroup
==========================

Creates a custom parameter group for an RDS database family. For more information about RDS parameter groups, see [Working with DB Parameter Groups](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Relational Database Service User Guide*.

This type can be declared in a template and referenced in the *`DBParameterGroupName`* parameter of [AWS::RDS::DBInstance](aws-properties-rds-database-instance.html "AWS::RDS::DBInstance").

Note

Applying a ParameterGroup to a DBInstance may require the instance to reboot, resulting in a database outage for the duration of the reboot.

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::RDS::DBParameterGroup",
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
A friendly description of the RDS parameter group. For example, `"My Parameter Group"`.

*Required*: Yes

*Type:* String

*Update requires*: Updates are not supported.

 `Family`   
The database family of this RDS parameter group. For example, `"MySQL5.1"`.

*Required*: Yes

*Type:* String

*Update requires*: Updates are not supported.

 `Parameters`   
The parameters to set for this RDS parameter group.

*Required*: No

*Type:* A JSON object consisting of string key-value pairs, as shown in the following example:

``` {.programlisting}
            
"Parameters" : {
   "Key1" : "Value1",
   "Key2" : "Value2",
   "Key3" : "Value3"
} 
          
```

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) or [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Changes to dynamic parameters are applied immediately. During an update, if you have static parameters (whether they were changed or not), triggers AWS CloudFormation to reboot the associated DB instance without failover.

 `Tags`   
The tags that you want to attach to the RDS parameter group.

*Required*: No

*Type*: A list of [resource tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
        { "Ref": "MyDBParameterGroup" }
      
```

For the RDS::DBParameterGroup with the logical ID "MyDBParameterGroup", `Ref` will return the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

