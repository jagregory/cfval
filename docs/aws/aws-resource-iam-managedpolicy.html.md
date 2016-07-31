AWS::IAM::ManagedPolicy
=======================

`AWS::IAM::ManagedPolicy` creates an AWS Identity and Access Management (IAM) managed policy for your AWS account that you can use to apply permissions to IAM users, groups, and roles. For more information about managed policies, see [Managed Policies and Inline Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/policies_managed-vs-inline.html) in the *IAM User Guide* guide.

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::IAM::ManagedPolicy",
  "Properties": {
    "Description" : String,
    "Groups" : [ String, ... ],
    "Path" : String,
    "PolicyDocument" : JSON object,
    "Roles" : [ String, ... ],
    "Users" : [ String, ... ]
  }
}
    
```

Properties
----------

 `Description`   
A description of the policy. For example, you can describe the permissions that are defined in the policy.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Groups`   
The names of groups to attach to this policy.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Path`   
The path for the policy. By default, the path is `/`. For more information, see [IAM Identifiers](http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide* guide.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `PolicyDocument`   
Policies that define the permissions for this managed policy. For more information about policy syntax, see [IAM Policy Elements Reference](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in *IAM User Guide*.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Roles`   
The names of roles to attach to this policy.

Note

If a policy has a `Ref` to a role and if a resource (such as `AWS::ECS::Service`) also has a `Ref` to the same role, add a `DependsOn` attribute to the resource so that the resource depends on the policy. This dependency ensures that the role's policy is available throughout the resource's lifecycle. For example, when you delete a stack with an `AWS::ECS::Service` resource, the `DependsOn` attribute ensures that the `AWS::ECS::Service` resource can complete its deletion before its role's policy is deleted.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Users`   
The names of users to attach to this policy.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the ARN.

In the following sample, the `Ref` function returns the ARN of the `CreateTestDBPolicy` managed policy, such as `arn:aws:iam::123456789012:policy/teststack-CreateTestDBPolicy-16M23YE3CS700`.

``` {.programlisting}
        { "Ref": "CreateTestDBPolicy" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following snippet creates a managed policy and associates it with the `TestDBGroup` group. The managed policy grants users permission to create t2.micro database instances. The database must use the MySQL database engine and the instance name must include the prefix `test`.

``` {.programlisting}
      "CreateTestDBPolicy" : {
  "Type" : "AWS::IAM::ManagedPolicy",
  "Properties" : {
    "Description" : "Policy for creating a test database",
    "Path" : "/",
    "PolicyDocument" :   {
      "Version":"2012-10-17", 
      "Statement" : [{
        "Effect" : "Allow",           
        "Action" : "rds:CreateDBInstance",
        "Resource" : {"Fn::Join" : [ "", [ "arn:aws:rds:", { "Ref" : "AWS::Region" }, ":", { "Ref" : "AWS::AccountId" }, ":db:test*" ] ]}, 
        "Condition" : {
          "StringEquals" : { "rds:DatabaseEngine" : "mysql" }
        }
      },
      {
        "Effect" : "Allow",           
        "Action" : "rds:CreateDBInstance",
        "Resource" : {"Fn::Join" : [ "", [ "arn:aws:rds:", { "Ref" : "AWS::Region" }, ":", { "Ref" : "AWS::Region" }, ":db:test*" ] ]}, 
        "Condition" : {
          "StringEquals" : { "rds:DatabaseClass" : "db.t2.micro" }
        }
      }]
    },
    "Groups" : ["TestDBGroup"]
  }
}
    
```
