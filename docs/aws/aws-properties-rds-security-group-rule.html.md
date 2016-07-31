Amazon RDS Security Group Rule
==============================

The Amazon RDS security group rule is an embedded property of the [AWS::RDS::DBSecurityGroup](aws-properties-rds-security-group.html "AWS::RDS::DBSecurityGroup") type.

Syntax
------

``` {.programlisting}
      
{
   "CIDRIP": String,
   "EC2SecurityGroupId": String,
   "EC2SecurityGroupName": String,
   "EC2SecurityGroupOwnerId": String
}     
    
```

Properties
----------

 `CIDRIP`   
The IP range to authorize.

For an overview of CIDR ranges, go to the [Wikipedia Tutorial](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

*Type*: String

*Required*: No

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EC2SecurityGroupId`   
Id of the VPC or EC2 Security Group to authorize.

For VPC DB Security Groups, use EC2SecurityGroupId. For EC2 Security Groups, use EC2SecurityGroupOwnerId and either EC2SecurityGroupName or EC2SecurityGroupId.

*Type*: String

*Required*: No

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EC2SecurityGroupName`   
Name of the EC2 Security Group to authorize.

For VPC DB Security Groups, use EC2SecurityGroupId. For EC2 Security Groups, use EC2SecurityGroupOwnerId and either EC2SecurityGroupName or EC2SecurityGroupId.

*Type*: String

*Required*: No

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `EC2SecurityGroupOwnerId`   
AWS Account Number of the owner of the EC2 Security Group specified in the EC2SecurityGroupName parameter. The AWS Access Key ID is not an acceptable value.

For VPC DB Security Groups, use EC2SecurityGroupId. For EC2 Security Groups, use EC2SecurityGroupOwnerId and either EC2SecurityGroupName or EC2SecurityGroupId.

*Type*: String

*Required*: No

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)


