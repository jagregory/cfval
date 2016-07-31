AWS::IAM::InstanceProfile
=========================

Creates an AWS Identity and Access Management (IAM) Instance Profile that can be used with IAM Roles for EC2 Instances.

For more information about IAM roles, see [Working with Roles](http://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html) in the *AWS Identity and Access Management User Guide*.

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::IAM::InstanceProfile",
   "Properties": {
      "Path": String,
      "Roles": [ IAM Roles ]
   }
}      
    
```

Properties
----------

 `Path`   
The path associated with this IAM instance profile. For information about IAM paths, see [Friendly Names and Paths](http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html#Identifiers_FriendlyNames) in the *AWS Identity and Access Management User Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Roles`   
The roles associated with this IAM instance profile.

*Required*: Yes

*Type*: List of references to AWS::IAM::Roles. Currently, a maximum of one role can be assigned to an instance profile.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
        { "Ref": "MyProfile" }
      
```

For the IAM::InstanceProfile with the logical ID "MyProfile", `Ref` will return the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
Returns the Amazon Resource Name (ARN) for the instance profile. For example:

``` {.programlisting}
              {"Fn::GetAtt" : ["MyProfile", "Arn"] }
            
```

This will return a value such as `“arn:aws:iam::1234567890:instance-profile/MyProfile-ASDNSDLKJ”`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

**Example IAM Role with Embedded Policy and Instance Profiles**

This example shows an embedded Policy in the IAM::Role. The policy is specified inline in the IAM::Role Policies property.

``` {.programlisting}
          {
   "AWSTemplateFormatVersion": "2010-09-09",
   "Resources": {
      "RootRole": {
         "Type": "AWS::IAM::Role",
         "Properties": {
            "AssumeRolePolicyDocument": {
               "Version" : "2012-10-17",
               "Statement": [ {
                  "Effect": "Allow",
                  "Principal": {
                     "Service": [ "ec2.amazonaws.com" ]
                  },
                  "Action": [ "sts:AssumeRole" ]
               } ]
            },
            "Path": "/",
            "Policies": [ {
               "PolicyName": "root",
               "PolicyDocument": {
                  "Version" : "2012-10-17",
                  "Statement": [ {
                     "Effect": "Allow",
                     "Action": "*",
                     "Resource": "*"
                  } ]
               }
               } ]
            }
      },
      "RootInstanceProfile": {
         "Type": "AWS::IAM::InstanceProfile",
         "Properties": {
            "Path": "/",
            "Roles": [ {
               "Ref": "RootRole"
            } ]
         }
      }
   }
}
        
```


