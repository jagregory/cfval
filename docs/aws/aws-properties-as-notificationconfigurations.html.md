Auto Scaling NotificationConfigurations
=======================================

The `NotificationConfigurations` property is an embedded property of the `AWS::AutoScaling::AutoScalingGroup` resource that specifies the events for which the Auto Scaling group sends notifications.

Syntax
------

``` {.programlisting}
      {
   "NotificationTypes" : [ String, ... ],
   "TopicARN" : String
}
    
```

Properties
----------

 `NotificationTypes`   
A list of event types that trigger a notification. Event types can include any of the following types: `autoscaling:EC2_INSTANCE_LAUNCH`, `autoscaling:EC2_INSTANCE_LAUNCH_ERROR`, `autoscaling:EC2_INSTANCE_TERMINATE`, `autoscaling:EC2_INSTANCE_TERMINATE_ERROR`, and `autoscaling:TEST_NOTIFICATION`. For more information about event types, see [DescribeAutoScalingNotificationTypes](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingNotificationTypes.html) in the *Auto Scaling API Reference*.

*Required*: Yes

*Type*: List of strings

 `TopicARN`   
The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic.

*Required*: Yes

*Type*: String

Examples
--------

For NotificationConfigurations snippets, see [Auto Scaling Group with Notifications](quickref-autoscaling.html#scenario-as-notification "Auto Scaling Group with Notifications").

