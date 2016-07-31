AWS::IAM::UserToGroupAddition
=============================

The AWS::IAM::UserToGroupAddition type adds AWS Identity and Access Management (IAM) users to a group.

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::IAM::UserToGroupAddition",
   "Properties": {
      "GroupName": String,
      "Users": [ User1, ... ]
   }
}     
    
```

Properties
----------

 `GroupName`   
The name of group to add users to.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Users`   
*Required*: Yes

*Type*: List of users

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "MyUserToGroupAddition" }
    
```

For the AWS::IAM::UserToGroupAddition with the logical ID "MyUserToGroupAddition", `Ref` will return the AWS resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Examples
-----------------

To view AWS::IAM::UserToGroupAddition snippets, see [Adding Users to a Group](quickref-iam.html#scenario-iam-addusertogroup "Adding Users to a Group").

