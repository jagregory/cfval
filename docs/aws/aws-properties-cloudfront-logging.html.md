CloudFront Logging
==================

`Logging` is a property of the [DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") property that enables Amazon CloudFront to deliver access logs for each distribution to an Amazon Simple Storage Service (S3) bucket.

Syntax
------

``` {.programlisting}
      {
  "Bucket" : String,
  "IncludeCookies" : Boolean,
  "Prefix" : String
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `Bucket`   
The Amazon S3 bucket address where access logs are stored, for example, `mybucket.s3.amazonaws.com`.

*Required*: Yes

*Type*: String

 `IncludeCookies`   
Indicates whether CloudFront includes cookies in access logs.

*Required*: No

*Type*: Boolean

 `Prefix`   
A prefix for the access log file names for this distribution.

*Required*: No

*Type*: String


