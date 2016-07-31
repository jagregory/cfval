AWS::Logs::MetricFilter
=======================

The `AWS::Logs::MetricFilter` resource creates a metric filter that describes how Amazon CloudWatch Logs extracts information from logs that you specify and transforms it into Amazon CloudWatch metrics. If you have multiple metric filters that are associated with a log group, all the filters are applied to the log streams in that group.

Syntax
------

``` {.programlisting}
      {
  "Type": "AWS::Logs::MetricFilter",    
  "Properties": {
    "FilterPattern": [String, ...],
    "LogGroupName": String,
    "MetricTransformations": [ MetricTransformations, ... ]
  }
}
    
```

Properties
----------

Note

For more information about constraints and values for each property, see [PutMetricFilter](http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutMetricFilter.html) in the *Amazon CloudWatch Logs API Reference*.

 `FilterPattern`   
Describes the pattern that CloudWatch Logs follows to interpret each entry in a log. For example, a log entry might contain fields such as timestamps, IP addresses, error codes, bytes transferred, and so on. You use the pattern to specify those fields and to specify what to look for in the log file. For example, if you're interested in error codes that begin with `1234`, your filter pattern might be `[timestamps, ip_addresses, error_codes = 1234*, size, ...]`.

*Required*: Yes

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LogGroupName`   
The name of an existing log group that you want to associate with this metric filter.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `MetricTransformations`   
Describes how to transform data from a log into a CloudWatch metric.

*Required*: Yes

*Type*: A list of [CloudWatch Logs MetricFilter MetricTransformation Property](aws-properties-logs-metricfilter-metrictransformation.html "CloudWatch Logs MetricFilter MetricTransformation Property")

Important

Currently, you can specify only one metric transformation for each metric filter. If you want to specify multiple metric transformations, you must specify multiple metric filters.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Examples
--------

### 

The following example sends a value of `1` to the `404Count` metric whenever the status code field includes a `404` value.

``` {.programlisting}
        "404MetricFilter": {
    "Type": "AWS::Logs::MetricFilter",
    "Properties": {
        "LogGroupName": { "Ref": "myLogGroup" },
        "FilterPattern": "[ip, identity, user_id, timestamp, request, status_code = 404, size]",
        "MetricTransformations": [
            {
                "MetricValue": "1",
                "MetricNamespace": "WebServer/404s",
                "MetricName": "404Count"
            }
        ]
    }
}
      
```

For an additional sample template, see [Amazon CloudWatch Logs Template Snippets](quickref-cloudwatchlogs.html "Amazon CloudWatch Logs Template Snippets").

