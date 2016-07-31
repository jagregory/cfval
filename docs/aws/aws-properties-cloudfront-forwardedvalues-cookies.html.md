CloudFront ForwardedValues Cookies
==================================

`Cookies` is a property of the [CloudFront ForwardedValues](aws-properties-cloudfront-forwardedvalues.html "CloudFront ForwardedValues") that describes which cookies are forwarded to the Amazon CloudFront origin.

Syntax
------

``` {.programlisting}
      {
  "Forward" : String,
  "WhitelistedNames" : [ String, ... ]
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `Forward`   
The cookies to forward to the origin of the cache behavior. You can specify `none`, `all`, or `whitelist`.

*Required*: Yes

*Type*: String

 `WhitelistedNames`   
The names of cookies to forward to the origin for the cache behavior.

*Required*: Conditional. Required if you specified `whitelist` for the `Forward` property.

*Type*: List of strings


