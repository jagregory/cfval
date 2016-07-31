Elastic Beanstalk Environment Tier Property Type
================================================

Describes the environment tier for an [AWS::ElasticBeanstalk::Environment](aws-properties-beanstalk-environment.html "AWS::ElasticBeanstalk::Environment") resource. For more information, see [Environment Tiers](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features-managing-env-tiers.html) in the *AWS Elastic Beanstalk Developer Guide*.

Syntax
------

``` {.programlisting}
      
{
   "Name" : String,
   "Type" : String,
   "Version" : String
}     
    
```

Members
-------

 `Name`   
The name of the environment tier. You can specify `WebServer` or `Worker`.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Type`   
The type of this environment tier. You can specify `Standard` for the `WebServer` tier or `SQS/HTTP` for the `Worker` tier.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Version`   
The version of this environment tier.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Example
-------

``` {.programlisting}
      "Tier" : {
  "Type" : "SQS/HTTP",
  "Name" : "Worker",
  "Version" : "1.0"
}
    
```
