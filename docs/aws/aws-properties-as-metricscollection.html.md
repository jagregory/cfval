Auto Scaling MetricsCollection
==============================

The `MetricsCollection` is a property of the [AWS::AutoScaling::AutoScalingGroup](aws-properties-as-group.html "AWS::AutoScaling::AutoScalingGroup") resource that describes the group metrics that an Auto Scaling group sends to CloudWatch. These metrics describe the group rather than any of its instances. For more information, see [EnableMetricsCollection](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_EnableMetricsCollection.html) in the *Auto Scaling API Reference*.

Syntax
------

``` {.programlisting}
      {
  "Granularity" : String,
  "Metrics" : [ String, ... ]
}
    
```

Properties
----------

 `Granularity`   
The frequency at which Auto Scaling sends aggregated data to CloudWatch. For example, you can specify `1Minute` to send aggregated data to CloudWatch every minute.

*Required*: Yes

*Type*: String

 `Metrics`   
The list of metrics to collect. If you don't specify any metrics, all metrics are enabled.

*Required*: No

*Type*: List of strings


