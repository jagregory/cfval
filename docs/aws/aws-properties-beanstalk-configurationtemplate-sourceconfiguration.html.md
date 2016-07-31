Elastic Beanstalk SourceConfiguration Property Type
===================================================

Use settings from another Elastic Beanstalk configuration template for the [AWS::ElasticBeanstalk::ConfigurationTemplate](aws-resource-beanstalk-configurationtemplate.html "AWS::ElasticBeanstalk::ConfigurationTemplate") resource type.

Syntax
------

``` {.programlisting}
      
{
   "ApplicationName" : String,
   "TemplateName" : String
}     
    
```

Members
-------

 `ApplicationName`   
The name of the Elastic Beanstalk application that contains the configuration template that you want to use.

*Required*: Yes

*Type*: String

 `TemplateName`   
The name of the configuration template.

*Required*: Yes

*Type*: String


