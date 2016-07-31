CloudFront DistributionConfig CacheBehavior
===========================================

`CacheBehavior` is a property of the [DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") property that describes the Amazon CloudFront (CloudFront) cache behavior when the requested URL matches a pattern.

Syntax
------

``` {.programlisting}
      {
  "AllowedMethods" : [ String, ... ],
  "CachedMethods" : [ String, ... ],
  "Compress" : Boolean,
  "DefaultTTL" : Number,
  "ForwardedValues" : ForwardedValues,
  "MaxTTL" : Number,
  "MinTTL" : Number,
  "PathPattern" : String,
  "SmoothStreaming" : Boolean,
  "TargetOriginId" : String,
  "TrustedSigners" : [ String, ... ],
  "ViewerProtocolPolicy" : String
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the corresponding element in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `AllowedMethods`   
HTTP methods that CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. You can specify `["HEAD", "GET"]`, `["GET", "HEAD",                      "OPTIONS"]`, or `["DELETE", "GET", "HEAD", "OPTIONS", "PATCH",                      "POST", "PUT"]`. If you don't specify a value, AWS CloudFormation specifies `["HEAD", "GET"]`.

*Required*: No

*Type*: List of strings

 `CachedMethods`   
HTTP methods for which CloudFront caches responses. You can specify `["HEAD",                      "GET"]` or `["GET", "HEAD", "OPTIONS"]`. If you don't specify a value, AWS CloudFormation specifies `["HEAD", "GET"]`.

*Required*: No

*Type*: List of strings

 `Compress`   
Indicates whether CloudFront automatically compresses certain files for this cache behavior. For more information, see [Serving Compressed Files](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide*.

*Required*: No

*Type*: Boolean

 `DefaultTTL`   
The default time in seconds that objects stay in CloudFront caches before CloudFront forwards another request to your custom origin to determine whether the object has been updated. This value applies only when your custom origin does not add HTTP headers, such as `Cache-Control max-age`, `Cache-Control                      s-maxage`, and `Expires` to objects.

By default, AWS CloudFormation specifies `86400` seconds (one day). If the value of the `MinTTL` property is greater than the default value, CloudFront uses the minimum Time to Live (TTL) value.

*Required*: No

*Type*: Number

 `ForwardedValues`   
Specifies how CloudFront handles query strings or cookies.

*Required*: Yes

*Type*: [ForwardedValues](aws-properties-cloudfront-forwardedvalues.html "CloudFront ForwardedValues") type

 `MaxTTL`   
The maximum time in seconds that objects stay in CloudFront caches before CloudFront forwards another request to your custom origin to determine whether the object has been updated. This value applies only when your custom origin does not add HTTP headers, such as `Cache-Control max-age`, `Cache-Control                      s-maxage`, and `Expires` to objects.

By default, AWS CloudFormation specifies `31536000` seconds (one year). If the value of the `MinTTL` or `DefaultTTL` property is greater than the maximum value, CloudFront uses the default TTL value.

*Required*: No

*Type*: Number

 `MinTTL`   
The minimum amount of time that you want objects to stay in the cache before CloudFront queries your origin to see whether the object has been updated.

*Required*: No

*Type*: Number

 `PathPattern`   
The pattern to which this cache behavior applies. For example, you can specify `images/*.jpg`.

When CloudFront receives an end-user request, CloudFront compares the requested path with path patterns in the order in which cache behaviors are listed in the template.

*Required*: Yes

*Type*: String

 `SmoothStreaming`   
Indicates whether to use the origin that is associated with this cache behavior to distribute media files in the Microsoft Smooth Streaming format. If you specify `true`, you can still use this cache behavior to distribute other content if the content matches the `PathPattern` value.

*Required*: No

*Type*: Boolean

 `TargetOriginId`   
The ID value of the origin to which you want CloudFront to route requests when a request matches the value of the `PathPattern` property.

*Required*: Yes

*Type*: String

 `TrustedSigners`   
A list of AWS accounts that can create signed URLs in order to access private content.

*Required*: No

*Type*: List of strings

 `ViewerProtocolPolicy`   
The protocol that users can use to access the files in the origin that you specified in the `TargetOriginId` property when a request matches the value of the `PathPattern` property. For more information about the valid values, see the `ViewerProtocolPolicy` elements in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

*Required*: Yes

*Type*: String


