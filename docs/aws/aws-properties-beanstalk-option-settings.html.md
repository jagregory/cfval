Elastic Beanstalk OptionSettings Property Type
==============================================

`OptionSettings` is an embedded property of the [AWS::ElasticBeanstalk::Environment](aws-properties-beanstalk-environment.html "AWS::ElasticBeanstalk::Environment") and [AWS::ElasticBeanstalk::ConfigurationTemplate](aws-resource-beanstalk-configurationtemplate.html "AWS::ElasticBeanstalk::ConfigurationTemplate") resources. You use the `OptionSettings` property to specify an array of options for the Elastic Beanstalk environment.

Syntax
------

``` {.programlisting}
      
{
   "Namespace" : String,
   "OptionName" : String,
   "Value" : String
}     
    
```

Members
-------

 `Namespace`   
A unique namespace identifying the option's associated AWS resource. For a list of namespaces that you can use, see [Configuration Options](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide*.

*Required*: Yes

*Type*: String

 `OptionName`   
The name of the configuration option. For a list of options that you can use, see [Configuration Options](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide*.

*Required*: Yes

*Type*: String

 `Value`   
The value of the setting.

*Required*: Yes

*Type*: String

Example
-------

This example of using `OptionSettings` is found in the AWS CloudFormation sample template: [ElasticBeanstalkSample.template](https://s3.amazonaws.com/cloudformation-templates-us-east-1/ElasticBeanstalkSample.template), which also provides an example of its use within an `AWS::ElasticBeanstalk::Application`.

``` {.programlisting}
      
"OptionSettings" : [ {
   "Namespace" : "aws:autoscaling:launchconfiguration",
   "OptionName" : "EC2KeyName",
   "Value" : { "Ref" : "KeyName" }
} ]   
    
```

See Also
--------

-   [ConfigurationOptionSetting](http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionSetting.html) in the *AWS Elastic Beanstalk Developer Guide*

-   [Option Values](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide*


