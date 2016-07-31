AWS::IAM::AccessKey
===================

The AWS::IAM::AccessKey resource type generates a secret access key and assigns it to an IAM user or AWS account.

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::IAM::AccessKey",
   "Properties": {
      "Serial": Integer,
      "Status": String,
      "UserName": String
   }
}     
    
```

Properties
----------

 `Serial`   
This value is specific to AWS CloudFormation and can only be *incremented*. Incrementing this value notifies AWS CloudFormation that you want to rotate your access key. When you update your stack, AWS CloudFormation will replace the existing access key with a new key.

*Required*: No

*Type*: Integer

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Status`   
The status of the access key. By default, AWS CloudFormation sets this property value to `Active`.

*Required*: No

*Type*: String

*Valid values:* `Active` or `Inactive`

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `UserName`   
The name of the user that the new key will belong to.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

Specifying this resource ID to the intrinsic `Ref` function will return the *`AccessKeyId`*. For example: `AKIAIOSFODNN7EXAMPLE`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `SecretAccessKey`   
Returns the secret access key for the specified AWS::IAM::AccessKey resource. For example: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYzEXAMPLEKEY`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

To view AWS::IAM::AccessKey snippets, see [Declaring an IAM Access Key Resource](quickref-iam.html#scenario-iam-accesskey "Declaring an IAM Access Key Resource").

