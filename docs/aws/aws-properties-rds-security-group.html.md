AWS::RDS::DBSecurityGroup
=========================

The AWS::RDS::DBSecurityGroup type is used to create or update an Amazon RDS DB Security Group. For more information about DB Security Groups, see [Working with DB Security Groups](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithSecurityGroups.html) in the *Amazon Relational Database Service Developer Guide*.

For details on the settings for DB security groups, see [CreateDBSecurityGroup](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBSecurityGroup.html).

When you specify an AWS::RDS::DBSecurityGroup as an argument to the `Ref` function, AWS CloudFormation returns the value of the *`DBSecurityGroupName`*.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::RDS::DBSecurityGroup",
   "Properties" :
   {
      "EC2VpcId" : { "Ref" : "myVPC" },
      "DBSecurityGroupIngress" : [ RDS Security Group Rule object 1, ... ],
      "GroupDescription" : String,
      "Tags" : [ Resource Tag, ... ]
   }
} 
    
```

Properties
----------

 `EC2VpcId`   
The Id of VPC. Indicates which VPC this DB Security Group should belong to.

*Type*: String

*Required*: Conditional. Must be specified to create a DB Security Group for a VPC; may not be specified otherwise.

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `DBSecurityGroupIngress`   
Network ingress authorization for an Amazon EC2 security group or an IP address range.

*Type*: List of [RDS Security Group Rules](aws-properties-rds-security-group-rule.html "Amazon RDS Security Group Rule").

*Required*: Yes

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `GroupDescription`   
Description of the security group.

*Type*: String

*Required*: Yes

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
The tags that you want to attach to the Amazon RDS DB security group.

*Required*: No

*Type*: A list of [resource tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type").

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Template Examples
-----------------

Tip

For more RDS template examples, see [Amazon RDS Template Snippets](quickref-rds.html "Amazon RDS Template Snippets").

### Single VPC security group

This template snippet creates/updates a single VPC security group, referred to by EC2SecurityGroupName.

``` {.programlisting}
        
"DBSecurityGroup": {
   "Type": "AWS::RDS::DBSecurityGroup",
   "Properties": {
      "EC2VpcId" : { "Ref" : "VpcId" },
      "DBSecurityGroupIngress": [
         {"EC2SecurityGroupName": { "Ref": "WebServerSecurityGroup"}}
      ],
      "GroupDescription": "Frontend Access"
   }
},
      
      
```

### Multiple VPC security groups

This template snippet creates/updates multiple VPC security groups.

``` {.programlisting}
        
{
   "Resources" : {
      "DBinstance" : {
         "Type" : "AWS::RDS::DBInstance",
         "Properties" : {
            "DBSecurityGroups" : [ {"Ref" : "DbSecurityByEC2SecurityGroup"} ],
            "AllocatedStorage" : "5",
            "DBInstanceClass" : "db.m1.small",
            "Engine" : "MySQL",
            "MasterUsername" : "YourName",
            "MasterUserPassword" : "YourPassword"
         },
         "DeletionPolicy" : "Snapshot"
      },
      "DbSecurityByEC2SecurityGroup" : {
         "Type" : "AWS::RDS::DBSecurityGroup",
         "Properties" : {
            "GroupDescription" : "Ingress for Amazon EC2 security group",
            "DBSecurityGroupIngress" : [ {
                  "EC2SecurityGroupId" : "sg-b0ff1111",
                  "EC2SecurityGroupOwnerId" : "111122223333"
               }, {
                  "EC2SecurityGroupId" : "sg-ffd722222",
                  "EC2SecurityGroupOwnerId" : "111122223333"
               } ]
         }
      }
   }
}
      
      
```
