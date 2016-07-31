AWS::IAM::Role
==============

Creates an AWS Identity and Access Management (IAM) role. Use an IAM role to enable applications running on an EC2 instance to securely access your AWS resources.

For more information about IAM roles, see [Working with Roles](http://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html) in the *AWS Identity and Access Management User Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::IAM::Role",
  "Properties": {
    "AssumeRolePolicyDocument": { JSON },
    "ManagedPolicyArns": [ String, ... ],
    "Path": String,
    "Policies": [ Policies, ... ],
    "RoleName": String
  }
}
    
```

Properties
----------

 `AssumeRolePolicyDocument`   
The trust policy that is associated with this role.

*Required*: Yes

*Type*: A JSON policy document

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Note

You can associate only one assume role policy with a role. For an example of an assume role policy, see [Template Examples](aws-resource-iam-role.html#cfn-iam-role-templateexamples "Template Examples").

 `ManagedPolicyArns`   
One or more managed policy ARNs to attach to this role.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Path`   
The path associated with this role. For information about IAM paths, see [Friendly Names and Paths](http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html#Identifiers_FriendlyNames) in *IAM User Guide*.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Policies`   
The policies to associate with this role. You can specify a policy inline or reference an external policy, such as a policy declared in an `AWS::IAM::Policy` or `AWS::IAM::ManagedPolicy` resource. For sample templates that demonstrates both embedded and external policies, see [Template Examples](aws-resource-iam-role.html#cfn-iam-role-templateexamples "Template Examples").

Important

The name of each policy for a role, user, or group must be unique. Duplicate policy names can cause IAM role updates to fail.

Note

If an external policy (such as `AWS::IAM::Policy` or `AWS::IAM::ManagedPolicy`) has a `Ref` to a role and if a resource (such as `AWS::ECS::Service`) also has a `Ref` to the same role, add a `DependsOn` attribute to the resource to make the resource depend on the external policy. This dependency ensures that the role's policy is available throughout the resource's lifecycle. For example, when you delete a stack with an `AWS::ECS::Service` resource, the `DependsOn` attribute ensures that AWS CloudFormation deletes the `AWS::ECS::Service` resource before deleting its role's policy.

*Required*: No

*Type*: List of [IAM Policies](aws-properties-iam-policy.html "IAM Policies")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `RoleName`   
A name for the IAM role. For valid values, see the `RoleName` parameter for the [`CreateRole`](http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *IAM API Reference*. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the group name.

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](using-iam-template.html#using-iam-capabilities "Acknowledging IAM Resources in AWS CloudFormation Templates").

Warning

Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple regions. To prevent this, we recommend using `Fn::Join` and `AWS::Region` to create a region-specific name, as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}`.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

### Notes on policies for IAM roles

For general information about IAM policies and policy documents, see [How to Write a Policy](http://docs.aws.amazon.com/IAM/latest/UserGuide/AccessPolicyLanguage_HowToWritePolicies.html) in *IAM User Guide*.

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
        { "Ref": "RootRole" }
      
```

For the IAM::Role with the logical ID "RootRole", `Ref` will return the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
Returns the Amazon Resource Name (ARN) for the instance profile. For example:

``` {.programlisting}
              {"Fn::GetAtt" : ["MyRole", "Arn"] }
            
```

This will return a value such as `“arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF”`.

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

**Example IAM Role with External Policy and Instance Profiles**

In this example, the Policy and InstanceProfile resources are specified externally to the IAM Role. They refer to the role by specifying its name, "RootRole", in their respective Roles properties.

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
            "Path": "/"
         }
      },
      "RolePolicies": {
         "Type": "AWS::IAM::Policy",
         "Properties": {
            "PolicyName": "root",
            "PolicyDocument": {
               "Version" : "2012-10-17",
               "Statement": [ {
                  "Effect": "Allow",
                  "Action": "*",
                  "Resource": "*"
               } ]
            },
            "Roles": [ {
               "Ref": "RootRole"
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

See Also
--------

-   [AWS Identity and Access Management Template Snippets](quickref-iam.html "AWS Identity and Access Management Template Snippets")

-   [AWS::IAM::InstanceProfile](aws-resource-iam-instanceprofile.html "AWS::IAM::InstanceProfile")


