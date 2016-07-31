AWS::Logs::LogStream
====================

The `AWS::Logs::LogStream` resource creates an Amazon CloudWatch Logs log stream in a log group. A log stream represents the sequence of events coming from an application instance or resource that you are monitoring. For more information, see [Monitoring Log Files](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatchLogs.html) in the *Amazon CloudWatch Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Logs::LogStream",
  "Properties" : {
    "LogGroupName" : String,
    "LogStreamName" : String
  }
}
    
```

Properties
----------

 `LogGroupName`   
The name of the log group where the log stream is created.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `LogStreamName`   
The name of the log stream to create. The name must be unique within the log group.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name, such as `MyAppLogStream`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

### 

The following example creates a CloudWatch Logs log stream named `MyAppLogStream` in the `exampleLogGroup` log group.

``` {.programlisting}
        "LogStream": {
  "Type": "AWS::Logs::LogStream",
  "Properties": {
    "LogGroupName" : "exampleLogGroup",
    "LogStreamName": "MyAppLogStream"
  }
}
      
```
