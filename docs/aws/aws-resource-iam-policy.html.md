AWS::IAM::Policy
================

The AWS::IAM::Policy resource associates an IAM policy with IAM users, roles, or groups. For more information about IAM policies, see [Overview of IAM Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) in the *IAM User Guide* guide.

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::IAM::Policy",
   "Properties": {
      "Groups" : [ String, ... ],
      "PolicyDocument" : JSON object,
      "PolicyName" : String,
      "Roles" : [ String, ... ],
      "Users" : [ String, ... ]
   }
}     
    
```

Properties
----------

 `Groups`   
The names of groups to which you want to add the policy.

*Required*: Conditional. You must specify at least one of the following properties: `Groups`, `Roles`, or `Users`.

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PolicyDocument`   
A policy document that contains permissions to add to the specified users or groups.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PolicyName`   
The name of the policy. If you specify multiple policies for an entity, specify unique names. For example, if you specify a list of policies for an IAM role, each policy must have a unique name.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Roles`   
The names of [AWS::IAM::Role](aws-resource-iam-role.html "AWS::IAM::Role")s to attach to this policy.

Note

If a policy has a `Ref` to a role and if a resource (such as `AWS::ECS::Service`) also has a `Ref` to the same role, add a `DependsOn` attribute to the resource so that the resource depends on the policy. This dependency ensures that the role's policy is available throughout the resource's lifecycle. For example, when you delete a stack with an `AWS::ECS::Service` resource, the `DependsOn` attribute ensures that the `AWS::ECS::Service` resource can complete its deletion before its role's policy is deleted.

*Required*: Conditional. You must specify at least one of the following properties: `Groups`, `Roles`, or `Users`.

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Users`   
The names of users for whom you want to add the policy.

*Required*: Conditional. You must specify at least one of the following properties: `Groups`, `Roles`, or `Users`.

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

### IAM Policy with policy group

``` {.programlisting}
        
{
   "Type" : "AWS::IAM::Policy",
   "Properties" : {
      "PolicyName" : "CFNUsers",
      "PolicyDocument" : {
         "Version" : "2012-10-17",
         "Statement": [ {
         "Effect"   : "Allow",
         "Action"   : [
            "cloudformation:Describe*",
            "cloudformation:List*",
            "cloudformation:Get*"
         ],
         "Resource" : "*"
         } ]
      },
      "Groups" : [ { "Ref" : "CFNUserGroup" } ]
   }
}        
      
```

### IAM Policy with specified role

``` {.programlisting}
        
{
   "Type": "AWS::IAM::Policy",
   "Properties": {
      "PolicyName": "root",
      "PolicyDocument": {
         "Version" : "2012-10-17",
         "Statement": [
            { "Effect": "Allow", "Action": "*", "Resource": "*" }
         ]
      },
      "Roles": [ { "Ref": "RootRole" } ]
   }
}        
      
```

To view more AWS::IAM::Policy snippets, see [Declaring an IAM Policy](quickref-iam.html#scenario-iam-policy "Declaring an IAM Policy").

