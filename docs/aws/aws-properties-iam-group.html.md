AWS::IAM::Group
===============

The AWS::IAM::Group resource creates an AWS Identity and Access Management (IAM) group.

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::IAM::Group",
  "Properties": {
    "GroupName": String,
    "ManagedPolicyArns": [ String, ... ],
    "Path": String,
    "Policies": [ Policies, ... ]
  }
}
    
```

Properties
----------

 `GroupName`   
A name for the IAM group. For valid values, see the `GroupName` parameter for the [`CreateGroup`](http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateGroup.html) action in the *IAM API Reference*. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the group name.

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](using-iam-template.html#using-iam-capabilities "Acknowledging IAM Resources in AWS CloudFormation Templates").

Warning

Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple regions. To prevent this, we recommend using `Fn::Join` and `AWS::Region` to create a region-specific name, as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}`.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ManagedPolicyArns`   
One or more managed policy ARNs to attach to this group.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Path`   
The path to the group. For more information about paths, see [IAM Identifiers](http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html) in the *IAM User Guide*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Policies`   
The policies to associate with this group. For information about policies, see [Overview of IAM Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?PoliciesOverview.html) in the *IAM User Guide*.

*Required*: No

*Type*: List of [IAM Policies](aws-properties-iam-policy.html "IAM Policies")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

Specifying this resource ID to the intrinsic `Ref` function will return the `GroupName`. For example: `mystack-mygroup-1DZETITOWEKVO`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
Returns the Amazon Resource Name (ARN) for the AWS::IAM::Group resource. For example: `arn:aws:iam::123456789012:group/mystack-mygroup-1DZETITOWEKVO`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

To view AWS::IAM::Group snippets, see [Declaring an IAM Group Resource](quickref-iam.html#scenario-iam-group "Declaring an IAM Group Resource")

