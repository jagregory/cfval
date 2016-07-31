CloudFront DistributionConfig CustomErrorResponse
=================================================

`CustomErrorResponse` is a property of the [CloudFront DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") resource that defines custom error messages for certain HTTP status codes.

Syntax
------

``` {.programlisting}
      {
  "ErrorCachingMinTTL" : Integer,
  "ErrorCode" : Integer,
  "ResponseCode" : Integer,
  "ResponsePagePath" : String
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `ErrorCachingMinTTL`   
The minimum amount of time, in seconds, that Amazon CloudFront caches the HTTP status code that you specified in the `ErrorCode` property. The default value is `300`.

*Required*: No

*Type*: Integer

 `ErrorCode`   
An HTTP status code for which you want to specify a custom error page. You can specify `400`, `403`, `404`, `405`, `414`, `500`, `501`, `502`, `503`, or `504`.

*Required*: Yes

*Type*: Integer

 `ResponseCode`   
The HTTP status code that CloudFront returns to viewer along with the custom error page. You can specify `200`, `400`, `403`, `404`, `405`, `414`, `500`, `501`, `502`, `503`, or `504`.

*Required*: Conditional. Required if you specified the `ResponsePagePath` property.

*Type*: Integer

 `ResponsePagePath`   
The path to the custom error page that CloudFront returns to a viewer when your origin returns the HTTP status code that you specified in the `ErrorCode` property. For example, you can specify `/404-errors/403-forbidden.html`.

*Required*: Conditional. Required if you specified the `ResponseCode` property.

*Type*: String


