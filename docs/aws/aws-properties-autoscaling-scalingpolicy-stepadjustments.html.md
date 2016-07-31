Auto Scaling ScalingPolicy StepAdjustments
==========================================

`StepAdjustments` is a property of the [AWS::AutoScaling::ScalingPolicy](aws-properties-as-policy.html "AWS::AutoScaling::ScalingPolicy") resource that describes a scaling adjustment based on the difference between the value of the aggregated CloudWatch metric and the breach threshold that you've defined for the alarm. For more information, see [StepAdjustment](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_StepAdjustment.html) in the *Auto Scaling API Reference*.

Syntax
------

``` {.programlisting}
      {
  "MetricIntervalLowerBound" : Number,
  "MetricIntervalUpperBound" : Number,
  "ScalingAdjustment" : Integer
}
    
```

Properties
----------

 `MetricIntervalLowerBound`   
The lower bound for the difference between the breach threshold and the CloudWatch metric. If the metric value exceeds the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.

*Required*: No

*Type*: Number

 `MetricIntervalUpperBound`   
The upper bound for the difference between the breach threshold and the CloudWatch metric. If the metric value exceeds the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.

*Required*: No

*Type*: Number

 `ScalingAdjustment`   
The amount by which to scale, based on the value that you specified in the `AdjustmentType` property. A positive value adds to the current capacity and a negative number subtracts from the current capacity.

*Required*: Yes

*Type*: Integer


