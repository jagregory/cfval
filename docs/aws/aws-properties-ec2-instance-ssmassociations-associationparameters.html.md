Amazon EC2 Instance SsmAssociations AssociationParameters
=========================================================

`AssociationParameters` is a property of the [Amazon EC2 Instance SsmAssociations](aws-properties-ec2-instance-ssmassociations.html "Amazon EC2 Instance SsmAssociations") property that specifies input parameter values for an Amazon EC2 Simple Systems Manager (SSM) document.

Syntax
------

``` {.programlisting}
      {
  "Key" : String,
  "Value" : [ String, ... ]
}
    
```

Properties
----------

 `Key`   
The name of an input parameter that is in the associated SSM document.

*Required*: Yes

*Type*: String

 `Value`   
The value of an input parameter.

*Required*: Yes

*Type*: List of strings


