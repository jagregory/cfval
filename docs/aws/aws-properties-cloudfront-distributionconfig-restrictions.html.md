CloudFront DistributionConfiguration Restrictions
=================================================

`Restrictions` is a property of the [CloudFront DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") property that lets you limit which viewers can access your content.

Syntax
------

``` {.programlisting}
      {
  "GeoRestriction" : GeoRestriction
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `GeoRestriction`   
The countries in which viewers are able to access your content.

*Required*: Yes

*Type*: [CloudFront DistributionConfig Restrictions GeoRestriction](aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html "CloudFront DistributionConfig Restrictions GeoRestriction")


