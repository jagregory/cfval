CloudFront DistributionConfig
=============================

DistributionConfig is a property of the [AWS::CloudFront::Distribution](aws-properties-cloudfront-distribution.html "AWS::CloudFront::Distribution") property that describes which Amazon CloudFront origin servers to get your files from when users request the files through your website or application.

Syntax
------

``` {.programlisting}
      
{
   "Aliases" : [ String, ... ],
   "CacheBehaviors" : [ CacheBehavior, ... ],
   "Comment" : String,
   "CustomErrorResponses" : [ CustomErrorResponse, ... ],
   "DefaultCacheBehavior" : DefaultCacheBehavior,
   "DefaultRootObject" : String,
   "Enabled" : Boolean,
   "Logging" : Logging,
   "Origins" : [ Origin, ... ],
   "PriceClass" : String,
   "Restrictions" : Restriction,
   "ViewerCertificate" : ViewerCertificate,
   "WebACLId" : String
}     
    
```

Properties
----------

 `Aliases`   
CNAMEs (alternate domain names), if any, for the distribution.

*Required*: No

*Type*: List of strings

 `CacheBehaviors`   
A list of CacheBehavior types for the distribution.

*Required*: No

*Type*: List of [CacheBehavior](aws-properties-cloudfront-cachebehavior.html "CloudFront DistributionConfig CacheBehavior")

 `Comment`   
Any comments that you want to include about the distribution.

*Required*: No

*Type*: String

 `CustomErrorResponses`   
Whether CloudFront replaces HTTP status codes in the `4xx` and `5xx` range with custom error messages before returning the response to the viewer.

*Required*: No

*Type* List of [CloudFront DistributionConfig CustomErrorResponse](aws-properties-cloudfront-distributionconfig-customerrorresponse.html "CloudFront DistributionConfig CustomErrorResponse")

 `DefaultCacheBehavior`   
The default cache behavior that is triggered if you do not specify the `CacheBehavior` property or if files don't match any of the values of `PathPattern` in the `CacheBehavior` property.

*Required*: Yes

*Type*: [DefaultCacheBehavior type](aws-properties-cloudfront-defaultcachebehavior.html "CloudFront DefaultCacheBehavior")

 `DefaultRootObject`   
The object (such as `index.html`) that you want CloudFront to request from your origin when the root URL for your distribution (such as `http://example.com/`) is requested.

Note

Specifying a default root object avoids exposing the contents of your distribution.

*Required*: No

*Type*: String

 `Enabled`   
Controls whether the distribution is enabled to accept end user requests for content.

*Required*: Yes

*Type*: Boolean

 `Logging`   
Controls whether access logs are written for the distribution. To turn on access logs, specify this property.

*Required*: No

*Type*: [Logging](aws-properties-cloudfront-logging.html "CloudFront Logging") type

 `Origins`   
A list of origins for this CloudFront distribution. For each origin, you can specify whether it is an Amazon S3 or custom origin.

*Required*: Yes

*Type*: List of [Origins](aws-properties-cloudfront-origin.html "CloudFront DistributionConfig Origin").

 `PriceClass`   
The price class that corresponds with the maximum price that you want to pay for the CloudFront service. For more information, see [Choosing the Price Class](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide*.

For more information about the constraints and valid values, see the `PriceClass` element for the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) data type in the *Amazon CloudFront API Reference*.

*Required*: No

*Type*: String

 `Restrictions`   
Specifies restrictions on who or how viewers can access your content.

*Required*: No

*Type*: [CloudFront DistributionConfiguration Restrictions](aws-properties-cloudfront-distributionconfig-restrictions.html "CloudFront DistributionConfiguration Restrictions")

 `ViewerCertificate`   
The certificate to use when viewers use HTTPS to request objects.

*Required*: No

*Type*: [CloudFront DistributionConfiguration ViewerCertificate](aws-properties-cloudfront-distributionconfig-viewercertificate.html "CloudFront DistributionConfiguration ViewerCertificate")

 `WebACLId`   
The AWS WAF [web ACL](aws-resource-waf-webacl.html "AWS::WAF::WebACL") to associate with this distribution. AWS WAF is a web application firewall that enables you to monitor the HTTP and HTTPS requests that are forwarded to CloudFront and to control who can access your content. CloudFront permits or forbids requests based on conditions that you specify, such as the IP addresses from which requests originate or the values of query strings.

*Required*: No

*Type*: String

See Also
--------

-   [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html) in the *Amazon CloudFront API Reference*


