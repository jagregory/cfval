IAM User LoginProfile
=====================

`LoginProfile` is a property of the [AWS::IAM::User](aws-properties-iam-user.html "AWS::IAM::User") resource that creates a login profile for users so that they can access the AWS Management Console.

Syntax
------

``` {.programlisting}
      {
  "Password" : String,
  "PasswordResetRequired" : Boolean
}
    
```

Properties
----------

 `Password`   
The password for the user.

*Required*: Yes

*Type*: String

 `PasswordResetRequired`   
Specifies whether the user is required to set a new password the next time the user logs in to the AWS Management Console.

*Required*: No

*Type*: Boolean


