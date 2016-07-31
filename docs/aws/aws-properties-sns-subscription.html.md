Amazon SNS Subscription Property Type
=====================================

`Subscription` is an embedded property of the `AWS::SNS::Topic` resource that describes the subscription endpoints for an Amazon Simple Notification Service (Amazon SNS) topic.

Syntax
------

``` {.programlisting}
      {
   "Endpoint" : String,
   "Protocol" : String
}
    
```

Properties
----------

 `Endpoint`   
The subscription's endpoint (format depends on the protocol). For more information, see the [Subscribe Endpoint](http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) parameter in the *Amazon Simple Notification Service API Reference*.

*Required*: Yes

*Type*: String

 `Protocol`   
The subscription's protocol. For more information, see the [Subscribe Protocol](http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) parameter in the *Amazon Simple Notification Service API Reference*.

*Required*: Yes

*Type*: String


