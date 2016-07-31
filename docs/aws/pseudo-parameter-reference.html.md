Pseudo Parameters Reference
===========================

Pseudo Parameters are parameters that are predefined by AWS CloudFormation. You do not declare them in your template. Use them the same way as you would a parameter, as the argument for the `Ref` function.

For example, the following fragment assigns the value of the *`AWS::Region`* pseudo parameter:

``` {.programlisting}
    
"Outputs" : {
   "MyStacksRegion" : { "Value" : { "Ref" : "AWS::Region" } }
}
  
```

The currently available pseudo parameters are listed here.

 AWS::AccountId   
Returns the AWS account ID of the account in which the stack is being created, such as `123456789012`.

 AWS::NotificationARNs   
Returns the list of notification Amazon Resource Names (ARNs) for the current stack.

For example:

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "MyNestedStack" : {
         "Type" : "AWS::CloudFormation::Stack",
         "Properties" : {
            "TemplateURL" : "https://my-website.com/stack-spec.json",
            "NotificationARNs" : {"Ref" : "AWS::NotificationARNs"}
         }
      }
   }
}
        
```

To get a single ARN from the list, use [Fn::Select](intrinsic-function-reference-select.html "Fn::Select"):

``` {.programlisting}
          
"myASGrpOne" : {
   "Type" : "AWS::AutoScaling::AutoScalingGroup",
   "Version" : "2009-05-15",
   "Properties" : {
      "AvailabilityZones" : [ "us-east-1a" ],
      "LaunchConfigurationName" : { "Ref" : "MyLaunchConfiguration" },
      "MinSize" : "0",
      "MaxSize" : "0",
      "NotificationConfigurations" : [{
         "TopicARN" : { "Fn::Select" : [ "0", { "Ref" : "AWS::NotificationARNs" } ] },
         "NotificationTypes" : [ "autoscaling:EC2_INSTANCE_LAUNCH", "autoscaling:EC2_INSTANCE_LAUNCH_ERROR" ]
      }]
   }
}
        
```

 AWS::NoValue   
Removes the corresponding resource property when specified as a return value in the `Fn::If` intrinsic function. For example, you can use the `AWS::NoValue` parameter when you want to use a snapshot for an Amazon RDS DB instance only if a snapshot ID is provided, as shown in the following snippet:

``` {.programlisting}
          "MyDB" : {
  "Type" : "AWS::RDS::DBInstance",
  "Properties" : {
    "AllocatedStorage" : "5",
    "DBInstanceClass" : "db.m1.small",
    "Engine" : "MySQL",
    "EngineVersion" : "5.5",
    "MasterUsername" : { "Ref" : "DBUser" },
    "MasterUserPassword" : { "Ref" : "DBPassword" },
    "DBParameterGroupName" : { "Ref" : "MyRDSParamGroup" },
    "DBSnapshotIdentifier" : {
      "Fn::If" : [
        "UseDBSnapshot",
        {"Ref" : "DBSnapshotName"},
        {"Ref" : "AWS::NoValue"}
      ]
    }
  }
}
        
```

If the `UseDBSnapshot` condition evaluates to true, AWS CloudFormation uses the `DBSnapshotName` parameter value for the `DBSnapshotIdentifier` property. If the condition evaluates to false, AWS CloudFormation removes the `DBSnapshotIdentifier` property.

 AWS::Region   
Returns a string representing the AWS Region in which the encompassing resource is being created, such as `us-west-2`.

 AWS::StackId   
Returns the ID of the stack as specified with the `aws cloudformation create-stack` command, such as `arn:aws:cloudformation:us-west-2:123456789012:stack/teststack/51af3dc0-da77-11e4-872e-1234567db123`.

 AWS::StackName   
Returns the name of the stack as specified with the `aws cloudformation create-stack` command, such as `teststack`.


