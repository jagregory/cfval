Amazon RDS OptionGroup OptionConfigurations
===========================================

Use the `OptionConfigurations` property to configure an option and its settings for an [AWS::RDS::OptionGroup](aws-resource-rds-optiongroup.html "AWS::RDS::OptionGroup") resource.

Syntax
------

``` {.programlisting}
      {
  "DBSecurityGroupMemberships" : [ String, ... ],
  "OptionName" : String,
  "OptionSettings" : [ OptionSettings, ... ],
  "Port" : Integer,
  "VpcSecurityGroupMemberships" : [ String, ... ]
}
    
```

Properties
----------

 `DBSecurityGroupMemberships`   
A list of database security group names for this option. If the option requires access to a port, the security groups must allow access to that port. If you specify this property, don't specify the `VPCSecurityGroupMemberships` property.

*Required*: No

*Type*: List of strings

 `OptionName`   
The name of the option. For more information about options, see [Working with Option Groups](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html) in the *Amazon Relational Database Service User Guide*.

*Required*: Yes

*Type*: String

 `OptionSettings`   
The settings for this option.

*Required*: No

*Type*: [Amazon RDS OptionGroup OptionConfigurations OptionSettings](aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html "Amazon RDS OptionGroup OptionConfigurations OptionSettings")

 `Port`   
The port number that this option uses.

*Required*: No

*Type*: Integer

 `VpcSecurityGroupMemberships`   
A list of VPC security group IDs for this option. If the option requires access to a port, the security groups must allow access to that port. If you specify this property, don't specify the `DBSecurityGroupMemberships` property.

*Required*: No

*Type*: List of strings


