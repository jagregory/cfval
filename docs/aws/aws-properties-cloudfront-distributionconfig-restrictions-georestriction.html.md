CloudFront DistributionConfig Restrictions GeoRestriction
=========================================================

`GeoRestriction` is a property of the [CloudFront DistributionConfiguration Restrictions](aws-properties-cloudfront-distributionconfig-restrictions.html "CloudFront DistributionConfiguration Restrictions") property that describes the countries in which Amazon CloudFront allows viewers to access your content.

Syntax
------

``` {.programlisting}
      {
  "Locations" : [ String, ... ],
  "RestrictionType" : String
}
    
```

Properties
----------

Note

For more information about the constraints and valid values of each property, see the elements table in the [DistributionConfig Complex Type](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/DistributionConfigDatatype.html#DistributionConfigDatatype_Elements) topic in the *Amazon CloudFront API Reference*.

 `Locations`   
The two-letter, uppercase country code for a country that you want to include in your blacklist or whitelist.

*Required*: Conditional. Required if you specified `blacklist` or `whitelist` for the `RestrictionType` property.

*Type*: List of strings

 `RestrictionType`   
The method to restrict distribution of your content:

 `blacklist`   
Prevents viewers in the countries that you specified from accessing your content.

 `whitelist`   
Allows viewers in the countries that you specified to access your content.

 `none`   
No distribution restrictions by country.

*Required*: Yes

*Type*: String


