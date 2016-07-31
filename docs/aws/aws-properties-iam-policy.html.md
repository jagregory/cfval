IAM Policies
============

Policies is a property of the [AWS::IAM::Role](aws-resource-iam-role.html "AWS::IAM::Role"), [AWS::IAM::Group](aws-properties-iam-group.html "AWS::IAM::Group"), and [AWS::IAM::User](aws-properties-iam-user.html "AWS::IAM::User") resources. The Policies property describes what actions are allowed on what resources. For more information about IAM policies, see [Overview of Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/PoliciesOverview.html) in *IAM User Guide*.

Syntax
------

``` {.programlisting}
      {
  "PolicyDocument" : JSON,
  "PolicyName" : String
}
    
```

Properties
----------

 `PolicyDocument`   
A policy document that describes what actions are allowed on which resources.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PolicyName`   
The name of the policy.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)


