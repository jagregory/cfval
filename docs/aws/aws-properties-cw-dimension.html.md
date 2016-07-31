CloudWatch Metric Dimension Property Type
=========================================

The Metric Dimension is an embedded property of the [AWS::CloudWatch::Alarm](aws-properties-cw-alarm.html "AWS::CloudWatch::Alarm") type. Dimensions are arbitrary name/value pairs that can be associated with a CloudWatch metric. You can specify a maximum of 10 dimensions for a given metric.

Syntax
------

``` {.programlisting}
      
{
   "Name" : String,
   "Value" : String
}     
    
```

Properties
----------

 `Name`   
The name of the dimension, from 1–255 characters in length.

*Required*: Yes

*Type*: String

 `Value`   
The value representing the dimension measurement, from 1–255 characters in length.

*Required*: Yes

*Type*: String

Examples
--------

### Two CloudWatch alarms with dimension values supplied by the Ref function

The [Ref](intrinsic-function-reference-ref.html "Ref") and [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt") intrinsic functions are often used to supply values for CloudWatch metric dimensions. Here is an example using the `Ref` function.

``` {.programlisting}
        
"CPUAlarmHigh": {
   "Type": "AWS::CloudWatch::Alarm",
   "Properties": {
      "AlarmDescription": "Scale-up if CPU is greater than 90% for 10 minutes",
      "MetricName": "CPUUtilization",
      "Namespace": "AWS/EC2",
      "Statistic": "Average",
      "Period": "300",
      "EvaluationPeriods": "2",
      "Threshold": "90",
      "AlarmActions": [ { "Ref": "WebServerScaleUpPolicy" } ],
      "Dimensions": [
         {
            "Name": "AutoScalingGroupName",
            "Value": { "Ref": "WebServerGroup" }
         }
      ],
      "ComparisonOperator": "GreaterThanThreshold"
   }
},
"CPUAlarmLow": {
   "Type": "AWS::CloudWatch::Alarm",
   "Properties": {
      "AlarmDescription": "Scale-down if CPU is less than 70% for 10 minutes",
      "MetricName": "CPUUtilization",
      "Namespace": "AWS/EC2",
      "Statistic": "Average",
      "Period": "300",
      "EvaluationPeriods": "2",
      "Threshold": "70",
      "AlarmActions": [ { "Ref": "WebServerScaleDownPolicy" } ],
      "Dimensions": [
         {
            "Name": "AutoScalingGroupName",
            "Value": { "Ref": "WebServerGroup" }
         }
      ],
      "ComparisonOperator": "LessThanThreshold"
   }
}        
      
```

See Also
--------

-   [Dimension](http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Dimension.html) in the *Amazon CloudWatch API Reference*

-   [Amazon CloudWatch Metrics, Namespaces, and Dimensions Reference](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html) in the *Amazon CloudWatch Developer Guide*


