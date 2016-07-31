AWS::CloudFront::Distribution
=============================

Creates an Amazon CloudFront web distribution. For general information about CloudFront distributions, see the [Introduction to Amazon CloudFront](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html) in the *Amazon CloudFront Developer Guide*. For specific information about creating CloudFront web distributions, see [POST Distribution](http://docs.aws.amazon.com/AmazonCloudFront/latest/APIReference/CreateDistribution.html) in the *Amazon CloudFront API Reference*.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::CloudFront::Distribution",
   "Properties" : {
      "DistributionConfig" : DistributionConfig
   }
}
      
    
```

Properties
----------

 `DistributionConfig`   
The distribution's configuration information.

*Required*: Yes

*Type*: [DistributionConfig](aws-properties-cloudfront-distributionconfig.html "CloudFront DistributionConfig") type

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

*Returns*: The CloudFront distribution ID. For example: `E27LVI50CSW06W`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `DomainName`   
*Returns*: The domain name of the resource. For example: `d2fadu0nynjpfn.cloudfront.net`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

To view AWS::CloudFront::Distribution snippets, see [Amazon CloudFront Template Snippets](quickref-cloudfront.html "Amazon CloudFront Template Snippets").

