AWS::CloudWatch::Alarm
======================

The AWS::CloudWatch::Alarm type creates an CloudWatch alarm.

This type supports updates. For more information about updating this resource, see [PutMetricAlarm](http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html). For more information about updating stacks, see [AWS CloudFormation Stacks Updates](using-cfn-updating-stacks.html "AWS CloudFormation Stacks Updates").

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::CloudWatch::Alarm",
   "Properties" : {
      "ActionsEnabled" : Boolean,
      "AlarmActions" : [ String, ... ],
      "AlarmDescription" : String,
      "AlarmName" : String,
      "ComparisonOperator" : String,
      "Dimensions" : [ Metric dimension, ... ],
      "EvaluationPeriods" : String,
      "InsufficientDataActions" : [ String, ... ],
      "MetricName" : String,
      "Namespace" : String,
      "OKActions" : [ String, ... ],
      "Period" : String,
      "Statistic" : String,
      "Threshold" : String,
      "Unit" : String
   }
}
      
    
```

Properties
----------

 `ActionsEnabled`   
Indicates whether or not actions should be executed during any changes to the alarm's state.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AlarmActions`   
The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Number (ARN). For more information about creating alarms and the actions you can specify, see [Creating Amazon CloudWatch Alarms](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/AlarmThatSendsEmail.html) in the *Amazon CloudWatch Developer Guide*.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AlarmDescription`   
The description for the alarm.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AlarmName`   
A name for the alarm. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the alarm name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `ComparisonOperator`   
The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand.

You can specify the following values: *`GreaterThanOrEqualToThreshold`* | *`GreaterThanThreshold`* | *`LessThanThreshold`* | *`LessThanOrEqualToThreshold`*

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Dimensions`   
The dimensions for the alarm's associated metric.

*Required*: No

*Type*: List of [Metric Dimension](aws-properties-cw-dimension.html "CloudWatch Metric Dimension Property Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EvaluationPeriods`   
The number of periods over which data is compared to the specified threshold.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `InsufficientDataActions`   
The list of actions to execute when this alarm transitions into an INSUFFICIENT\_DATA state from any other state. Each action is specified as an Amazon Resource Number (ARN). Currently the only action supported is publishing to an Amazon SNS topic or an Amazon Auto Scaling policy.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MetricName`   
The name for the alarm's associated metric. For more information about the metrics that you can specify, see [Amazon CloudWatch Namespaces, Dimensions, and Metrics Reference](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html) in the *Amazon CloudWatch Developer Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Namespace`   
The namespace for the alarm's associated metric.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `OKActions`   
The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Number (ARN). Currently the only action supported is publishing to an Amazon SNS topic or an Amazon Auto Scaling policy.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Period`   
The time over which the specified statistic is applied. You must specify a time in seconds that is also a multiple of 60.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Statistic`   
The statistic to apply to the alarm's associated metric.

You can specify the following values: `SampleCount` | `Average` | `Sum` | `Minimum` | `Maximum`

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Threshold`   
The value against which the specified statistic is compared.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Unit`   
The unit for the alarm's associated metric.

You can specify the following values: Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When you specify an AWS::CloudWatch::Alarm type as an argument to the `Ref` function, AWS CloudFormation returns the value of the *`AlarmName`*.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Examples
--------

For sample template snippets, see [Amazon CloudWatch Template Snippets](quickref-cloudwatch.html "Amazon CloudWatch Template Snippets").

