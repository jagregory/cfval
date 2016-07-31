AWS::IAM::User
==============

The AWS::IAM::User resource creates a user.

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::IAM::User",
  "Properties": {
    "Groups": [ String, ... ],
    "LoginProfile": LoginProfile Type,
    "ManagedPolicyArns": [ String, ... ],
    "Path": String,
    "Policies": [ Policies, ... ],
    "UserName": String
  }
}
    
```

Properties
----------

 `Groups`   
A name of a group to which you want to add the user.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LoginProfile`   
Creates a login profile so that the user can access the AWS Management Console.

*Required*: No

*Type*: [IAM User LoginProfile](aws-properties-iam-user-loginprofile.html "IAM User LoginProfile")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `ManagedPolicyArns`   
One or more managed policy ARNs to attach to this user.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Path`   
The path for the user name. For more information about paths, see [IAM Identifiers](http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html) in the *IAM User Guide*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Policies`   
The policies to associate with this user. For information about policies, see [Overview of IAM Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?PoliciesOverview.html) in the *IAM User Guide*.

Note

If you specify multiple polices, specify unique values for the policy name. If you don't, updates to the IAM user will fail.

*Required*: No

*Type*: List of [IAM Policies](aws-properties-iam-policy.html "IAM Policies")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `UserName`   
A name for the IAM user. For valid values, see the `UserName` parameter for the [`CreateUser`](http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateUser.html) action in the *IAM API Reference*. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the group name.

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](using-iam-template.html#using-iam-capabilities "Acknowledging IAM Resources in AWS CloudFormation Templates").

Warning

Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple regions. To prevent this, we recommend using `Fn::Join` and `AWS::Region` to create a region-specific name, as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}`.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

Specifying this resource ID to the intrinsic Ref function will return the `UserName`. For example: `mystack-myuser-1CCXAFG2H2U4D`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
Returns the Amazon Resource Name (ARN) for the specified AWS::IAM::User resource. For example: `arn:aws:iam::123456789012:user/mystack-myuser-1CCXAFG2H2U4D`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

To view AWS::IAM::User snippets, see: [Declaring an IAM User Resource](quickref-iam.html#scenario-iam-user "Declaring an IAM User Resource").

