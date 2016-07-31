AWS::KMS::Key
=============

The `AWS::KMS::Key` resource creates a customer master key (CMK) in AWS Key Management Service (AWS KMS). Users (customers) can use the master key to encrypt their data stored in AWS services that are integrated with AWS KMS or within their applications. For more information, see [What is the AWS Key Management Service?](http://docs.aws.amazon.com/kms/latest/developerguide/) in the *AWS Key Management Service Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::KMS::Key",
  "Properties" : {
    "Description" : String,
    "Enabled" : Boolean,
    "EnableKeyRotation" : Boolean,
    "KeyPolicy" : JSON object    
  }
}
    
```

Properties
----------

 `Description`   
A description of the key. Use a description that helps your users decide whether the key is appropriate for a particular task.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Enabled`   
Indicates whether the key is available for use. AWS CloudFormation sets this value to `true` by default.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EnableKeyRotation`   
Indicates whether AWS KMS rotates the key. AWS CloudFormation sets this value to `false` by default.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `KeyPolicy`   
An AWS Identity and Access Management (IAM) policy to attach to the key. Use a policy to specify who has permission to use the key and which actions they can perform. For more information, see [Key Policies](http://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the *AWS Key Management Service Developer Guide*.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, it returns the key ID, such as `123ab456-a4c2-44cb-95fd-b781f32fbb37`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

### 

The following example creates a custom CMK, which permits the IAM user `Alice` to administer the key and allows `Bob` to use the key for encrypting and decrypting data.

``` {.programlisting}
        "myKey" : {
  "Type" : "AWS::KMS::Key",
  "Properties" : {
    "Description" : "A sample key",
    "KeyPolicy" : {
      "Version": "2012-10-17",
      "Id": "key-default-1",
      "Statement": [
        {
          "Sid": "Allow administration of the key",
          "Effect": "Allow",
          "Principal": { "AWS": "arn:aws:iam::123456789012:user/Alice" },
          "Action": [
            "kms:Create*",
            "kms:Describe*",
            "kms:Enable*",
            "kms:List*",
            "kms:Put*",
            "kms:Update*",
            "kms:Revoke*",
            "kms:Disable*",
            "kms:Get*",
            "kms:Delete*",
            "kms:ScheduleKeyDeletion",
            "kms:CancelKeyDeletion"
          ],
          "Resource": "*"
        },
        {
          "Sid": "Allow use of the key",
          "Effect": "Allow",
          "Principal": { "AWS": "arn:aws:iam::123456789012:user/Bob" },
          "Action": [
            "kms:Encrypt",
            "kms:Decrypt",
            "kms:ReEncrypt",
            "kms:GenerateDataKey*",
            "kms:DescribeKey"
          ], 
          "Resource": "*"
        }    
      ]
    }
  }
}
      
```
