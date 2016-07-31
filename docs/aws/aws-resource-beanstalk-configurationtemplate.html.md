AWS::ElasticBeanstalk::ConfigurationTemplate
============================================

Creates a configuration template for an Elastic Beanstalk application. You can use configuration templates to deploy different versions of an application by using the configuration settings that you define in the configuration template.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ElasticBeanstalk::ConfigurationTemplate",
  "Properties" : {  
    "ApplicationName" : String,
    "Description" : String,
    "EnvironmentId" : String,
    "OptionSettings" : [ OptionSetting, ... ],
    "SolutionStackName" : String,
    "SourceConfiguration" : Source configuration
  } 
}
    
```

Members
-------

 `ApplicationName`   
Name of the Elastic Beanstalk application that is associated with this configuration template.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Description`   
An optional description for this configuration.

*Type*: String

*Required*: No

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `EnvironmentId`   
An environment whose settings you want to use to create the configuration template. You must specify this property if you don't specify the `SolutionStackName` or `SourceConfiguration` properties.

*Type*: String

*Required*: Conditional

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `OptionSettings`   
A list of [OptionSettings](aws-properties-beanstalk-option-settings.html "Elastic Beanstalk OptionSettings Property Type") for this Elastic Beanstalk configuration. For a complete list of Elastic Beanstalk configuration options, see [Option Values](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html), in the *AWS Elastic Beanstalk Developer Guide*.

*Type*: A list of [OptionSettings](aws-properties-beanstalk-option-settings.html "Elastic Beanstalk OptionSettings Property Type").

*Required*: No

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `SolutionStackName`   
The name of an Elastic Beanstalk solution stack that this configuration will use. A solution stack specifies the operating system, architecture, and application server for a configuration template, such as `64bit Amazon Linux 2013.09                      running Tomcat 7 Java 7`. For more information, see [Supported Platforms](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the *AWS Elastic Beanstalk Developer Guide*.

You must specify this property if you don't specify the `EnvironmentId` or `SourceConfiguration` properties.

*Type*: String

*Required*: Conditional

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SourceConfiguration`   
A configuration template that is associated with another Elastic Beanstalk application. If you specify the `SolutionStackName` property and the `SourceConfiguration` property, the solution stack in the source configuration template must match the value that you specified for the `SolutionStackName` property.

You must specify this property if you don't specify the `EnvironmentId` or `SolutionStackName` properties.

*Type*: [Elastic Beanstalk SourceConfiguration Property Type](aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html "Elastic Beanstalk SourceConfiguration Property Type")

*Required*: Conditional

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

This example of an ElasticBeanstalk ConfigurationTemplate is found in the AWS CloudFormation sample template [ElasticBeanstalkSample.template](https://s3.amazonaws.com/cloudformation-templates-us-east-1/ElasticBeanstalkSample.template), which also provides an example of its use within an AWS::ElasticBeanstalk::Application.

``` {.programlisting}
      "myConfigTemplate" : { 
  "Type" : "AWS::ElasticBeanstalk::ConfigurationTemplate",
  "Properties" : {
    "ApplicationName" :{"Ref" : "myApp"},
    "Description" : "my sample configuration template",
    "EnvironmentId" : "",
    "SourceConfiguration" : {
      "ApplicationName" : {"Ref" : "mySecondApp"},
      "TemplateName" : {"Ref" : "mySourceTemplate"}
    }, 
    "SolutionStackName" : "64bit Amazon Linux running PHP 5.3",
    "OptionSettings" : [ {
      "Namespace" : "aws:autoscaling:launchconfiguration",
      "OptionName" : "EC2KeyName",
      "Value" : { "Ref" : "KeyName" }
    } ]
  }
}
    
```

See Also
--------

-   [AWS::ElasticBeanstalk::Application](aws-properties-beanstalk.html "AWS::ElasticBeanstalk::Application")

-   [Option Values](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide*

-   For a complete Elastic Beanstalk sample template, see [Elastic Beanstalk Template Snippets](quickref-elasticbeanstalk.html "Elastic Beanstalk Template Snippets").


