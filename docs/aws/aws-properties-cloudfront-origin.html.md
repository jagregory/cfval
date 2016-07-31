CloudFront DistributionConfig Origin
====================================

`Origin` is a property of the [DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") property that describes an Amazon CloudFront distribution origin.

Syntax
------

``` {.programlisting}
      {
  "CustomOriginConfig" : Custom Origin,
  "DomainName" : String,
  "Id" : String,
  "OriginPath" : String,
  "S3OriginConfig" : S3 Origin
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `CustomOriginConfig`   
Origin information to specify a custom origin.

*Required*: Conditional. You cannot use `CustomOriginConfig` and `S3OriginConfig` in the same distribution, but you *must* specify one or the other.

*Type*: [CustomOrigin](aws-properties-cloudfront-customorigin.html "CloudFront DistributionConfig Origin CustomOrigin") type

 `DomainName`   
The DNS name of the Amazon Simple Storage Service (S3) bucket or the HTTP server from which you want CloudFront to get objects for this origin.

*Required*: Yes

*Type*: String

 `Id`   
An identifier for the origin. The value of `Id` must be unique within the distribution.

*Required*: Yes

*Type*: String

 `OriginPath`   
The path that CloudFront uses to request content from an S3 bucket or custom origin. The combination of the `DomainName` and `OriginPath` properties must resolve to a valid path. The value must start with a slash mark (`/`) and cannot end with a slash mark.

*Required*: No

*Type*: String

 `S3OriginConfig`   
Origin information to specify an S3 origin.

*Required*: Conditional. You cannot use `S3OriginConfig` and `CustomOriginConfig` in the same distribution, but you *must* specify one or the other.

*Type*: [S3Origin](aws-properties-cloudfront-s3origin.html "CloudFront DistributionConfig Origin S3Origin") type


