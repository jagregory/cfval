Elastic Load Balancing AccessLoggingPolicy
==========================================

The `AccessLoggingPolicy` property describes where and how access logs are stored for the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") resource.

Syntax
------

``` {.programlisting}
      {
  "EmitInterval" : Integer,
  "Enabled" : Boolean,
  "S3BucketName" : String,
  "S3BucketPrefix" : String
}
    
```

Properties
----------

 `EmitInterval`   
The interval for publishing access logs in minutes. You can specify an interval of either 5 minutes or 60 minutes.

*Required*: No

*Type*: Integer

 `Enabled`   
Whether logging is enabled for the load balancer.

*Required*: Yes

*Type*: Boolean

 `S3BucketName`   
The name of an Amazon S3 bucket where access log files are stored.

*Required*: Yes

*Type*: String

 `S3BucketPrefix`   
A prefix for the all log object keys, such as `my-load-balancer-logs/prod`. If you store log files from multiple sources in a single bucket, you can use a prefix to distinguish each log file and its source.

*Required*: No

*Type*: String


