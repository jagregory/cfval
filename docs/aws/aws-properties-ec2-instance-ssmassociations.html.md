Amazon EC2 Instance SsmAssociations
===================================

`SsmAssociations` is a property of the [AWS::EC2::Instance](aws-properties-ec2-instance.html "AWS::EC2::Instance") resource that specifies the Amazon EC2 Simple Systems Manager (SSM) document and parameter values to associate with an instance.

Syntax
------

``` {.programlisting}
      {
  "AssociationParameters" : [ Parameters, ... ],
  "DocumentName" : String
}
    
```

Properties
----------

 `AssociationParameters`   
The input parameter values to use with the associated SSM document.

*Required*: No

*Type*: List of [Amazon EC2 Instance SsmAssociations AssociationParameters](aws-properties-ec2-instance-ssmassociations-associationparameters.html "Amazon EC2 Instance SsmAssociations AssociationParameters")

 `DocumentName`   
The name of an SSM document to associate with the instance.

*Required*: Yes

*Type*: String


