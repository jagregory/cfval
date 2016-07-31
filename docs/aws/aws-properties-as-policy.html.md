AWS::AutoScaling::ScalingPolicy
===============================

The AWS::AutoScaling::ScalingPolicy resource adds a scaling policy to an auto scaling group. A scaling policy specifies whether to scale the auto scaling group up or down, and by how much. For more information on scaling policies, see [Scaling by Policy](http://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/scaling_plan.html#scaling_policies) in the Auto Scaling Developer Guide.

You can use a scaling policy together with an CloudWatch alarm. An CloudWatch alarm can automatically initiate actions on your behalf, based on parameters you specify. A scaling policy is one type of action that an alarm can initiate. For a snippet showing how to create an Auto Scaling policy that is triggered by an CloudWatch alarm, see [Auto Scaling Policy Triggered by CloudWatch Alarm](quickref-autoscaling.html#scenario-as-policy "Auto Scaling Policy Triggered by CloudWatch Alarm").

This type supports updates. For more information about updating this resource, see [PutScalingPolicy](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScalingPolicy.html).

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::AutoScaling::ScalingPolicy",
   "Properties" : {
      "AdjustmentType" : String,
      "AutoScalingGroupName" : String,
      "Cooldown" : String,
      "EstimatedInstanceWarmup" : Integer,
      "MetricAggregationType" : String,
      "MinAdjustmentMagnitude" : Integer,
      "PolicyType" : String,
      "ScalingAdjustment" : Integer,
      "StepAdjustments" : [ StepAdjustments, ... ]
   }
}      
    
```

Properties
----------

 `AdjustmentType`   
Specifies whether the *`ScalingAdjustment`* is an absolute number or a percentage of the current capacity. Valid values are *`ChangeInCapacity`*, *`ExactCapacity`*, and *`PercentChangeInCapacity`*.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AutoScalingGroupName`   
The name or Amazon Resource Name (ARN) of the Auto Scaling Group that you want to attach the policy to.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Cooldown`   
The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.

Do not specify this property if you are using the `StepScaling` policy type.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EstimatedInstanceWarmup`   
The estimated time, in seconds, until a newly launched instance can send metrics to CloudWatch. By default, Auto Scaling uses the cooldown period, as specified in the `Cooldown` property.

Do not specify this property if you are using the `SimpleScaling` policy type.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MetricAggregationType`   
The aggregation type for the CloudWatch metrics. You can specify `Minimum`, `Maximum`, or `Average`. By default, AWS CloudFormation specifies `Average`.

Do not specify this property if you are using the `SimpleScaling` policy type.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MinAdjustmentMagnitude`   
For the `PercentChangeInCapacity` adjustment type, the minimum number of instances to scale. The scaling policy changes the desired capacity of the Auto Scaling group by a minimum of this many instances. This property replaces the `MinAdjustmentStep` property.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `PolicyType`   
An Auto Scaling policy type. You can specify `SimpleScaling` or `StepScaling`. By default, AWS CloudFormation specifies `SimpleScaling`. For more information, see [Scaling Policy Types](http://docs.aws.amazon.com/autoscaling/latest/userguide/as-scale-based-on-demand.html#as-scaling-types) in the *Auto Scaling User Guide*.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `ScalingAdjustment`   
The number of instances by which to scale. The `AdjustmentType` property determines whether AWS CloudFormation interprets this number as an absolute number (when the *`ExactCapacity`*value is specified) or as a percentage of the existing Auto Scaling group size (when the *`PercentChangeInCapacity`* value is specified). A positive value adds to the current capacity and a negative value subtracts from the current capacity.

*Required*: Conditional. This property is required if the policy type is `SimpleScaling`. This property is not supported with any other policy type.

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `StepAdjustments`   
A set of adjustments that enable you to scale based on the size of the alarm breach.

*Required*: Conditional. This property is required if the policy type is `StepScaling`. This property is not supported with any other policy type.

*Type*: List of [Auto Scaling ScalingPolicy StepAdjustments](aws-properties-autoscaling-scalingpolicy-stepadjustments.html "Auto Scaling ScalingPolicy StepAdjustments")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

When you specify an AWS::AutoScaling::ScalingPolicy type as an argument to the `Ref` function, AWS CloudFormation returns the policy Amazon Resource Name (ARN), such as `arn:aws:autoscaling:us-east-1:123456789012:scalingPolicy:ab12c4d5-a1b2-a1b2-a1b2-ab12c4d56789:autoScalingGroupName/myStack-AutoScalingGroup-AB12C4D5E6:policyName/myStack-myScalingPolicy-AB12C4D5E6`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

### Simple policy type

The following example is a simple scaling policy that increases the number instances by one when it is triggered.

``` {.programlisting}
        "SimpleScaling" : {
  "Type" : "AWS::AutoScaling::ScalingPolicy",
  "Properties" : {
    "AdjustmentType" : "ExactCapacity",
    "PolicyType" : "SimpleScaling", 
    "Cooldown" : "60",
    "AutoScalingGroupName" : { "Ref" : "ASG" },
    "ScalingAdjustment" : 1
  }
}
      
```

### Step policy type

The following example is a step scaling policy that increases the number instances by one or two, depending on the size of the alarm breach. For a breach that is less than 50 units than the threshold value, the policy increases the number of instances by one. For a breach that is 50 units or more higher than the threshold, the policy increases the number of instances by two.

``` {.programlisting}
        "StepScaling" : {
  "Type" : "AWS::AutoScaling::ScalingPolicy",
  "Properties" : {
    "AdjustmentType" : "ExactCapacity",
    "AutoScalingGroupName" : { "Ref" : "ASG" },
    "PolicyType" : "StepScaling",
    "MetricAggregationType" : "Average",
    "EstimatedInstanceWarmup" : "60",
    "StepAdjustments": [
      {
        "MetricIntervalLowerBound": "0",
        "MetricIntervalUpperBound" : "50",
        "ScalingAdjustment": "1"
      },
      {
        "MetricIntervalLowerBound": "50",
        "ScalingAdjustment": "2"
      }
    ]
  }
}
      
```
