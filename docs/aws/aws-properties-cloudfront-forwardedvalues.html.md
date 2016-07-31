CloudFront ForwardedValues
==========================

`ForwardedValues` is a property of the [DefaultCacheBehavior](aws-properties-cloudfront-defaultcachebehavior.html "CloudFront DefaultCacheBehavior") and [CacheBehavior](aws-properties-cloudfront-cachebehavior.html "CloudFront DistributionConfig CacheBehavior") properties that indicates whether Amazon CloudFront forwards query strings or cookies.

Syntax
------

``` {.programlisting}
      {
  "Cookies" : Cookies,
  "Headers" : [ String, ... ],
  "QueryString" : Boolean
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `Cookies`   
Forwards specified cookies to the origin of the cache behavior. For more information, see [Configuring CloudFront to Cache Based on Cookies](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the *Amazon CloudFront Developer Guide*.

*Required*: No

*Type*: [CloudFront ForwardedValues Cookies](aws-properties-cloudfront-forwardedvalues-cookies.html "CloudFront ForwardedValues Cookies")

 `Headers`   
Specifies the headers that you want Amazon CloudFront to forward to the origin for this cache behavior (whitelisted headers). For the headers that you specify, Amazon CloudFront also caches separate versions of a specified object that is based on the header values in viewer requests.

For custom origins, if you specify a single asterisk (["\*"]), all headers are forwarded. If you don't specify a value, only the default headers are forwarded. For Amazon S3 origins, you can forward only selected headers; specifying \* is not supported. For more information, see [Configuring CloudFront to Cache Objects Based on Request Headers](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/header-caching.html) in the *Amazon CloudFront Developer Guide*.

*Required*: No

*Type*: List of strings

 `QueryString`   
Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior. If so, specify `true`; if not, specify `false`. For more information, see [Configuring CloudFront to Cache Based on Query String Parameters](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/QueryStringParameters.html) in the *Amazon CloudFront Developer Guide*.

*Required*: Yes

*Type*: Boolean


