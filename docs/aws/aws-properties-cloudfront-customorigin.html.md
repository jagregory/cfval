CloudFront DistributionConfig Origin CustomOrigin
=================================================

CustomOrigin is a property of the [Amazon CloudFront Origin](aws-properties-cloudfront-origin.html "CloudFront DistributionConfig Origin") property that describes an HTTP server.

Syntax
------

``` {.programlisting}
      
{
   "HTTPPort" : String,
   "HTTPSPort" : String,
   "OriginProtocolPolicy" : String
}     
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `HTTPPort`   
The HTTP port the custom origin listens on.

*Required*: No

*Type*: String

 `HTTPSPort`   
The HTTPS port the custom origin listens on.

*Required*: No

*Type*: String

 `OriginProtocolPolicy`   
The origin protocol policy to apply to your origin.

*Required*: Yes

*Type*: String


