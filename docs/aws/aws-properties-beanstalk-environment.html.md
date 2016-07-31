AWS::ElasticBeanstalk::Environment
==================================

Creates or updates an AWS Elastic Beanstalk environment.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::ElasticBeanstalk::Environment",
   "Properties" : {
      "ApplicationName" : String,
      "CNAMEPrefix" : String,
      "Description" :  String,
      "EnvironmentName" :  String,
      "OptionSettings" : [ OptionSettings, ... ],
      "SolutionStackName" : String,
      "Tags" : [ Resource Tag, ... ],
      "TemplateName" : String,
      "Tier" : Environment Tier,
      "VersionLabel" : String
   }
}
      
    
```

Properties
----------

 `ApplicationName`   
The name of the application that is associated with this environment.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `CNAMEPrefix`   
A prefix for your Elastic Beanstalk environment URL.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Description`   
A description that helps you identify this environment.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `EnvironmentName`   
A name for the Elastic Beanstalk environment. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the environment name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `OptionSettings`   
Key-value pairs defining configuration options for this environment. These options override the values that are defined in the solution stack or the configuration template. If you remove any options during a stack update, the removed options revert to default values.

*Required*: No

*Type*: A list of [OptionSettings](aws-properties-beanstalk-option-settings.html "Elastic Beanstalk OptionSettings Property Type").

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `SolutionStackName`   
The name of an Elastic Beanstalk solution stack that this configuration will use. For more information, see [Supported Platforms](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the *AWS Elastic Beanstalk Developer Guide*. You must specify either this parameter or an Elastic Beanstalk configuration template name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Tags`   
An arbitrary set of tags (key–value pairs) for this environment.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: You can update tags only if you update another property that requires that the environment be replaced, such as the `ApplicationName` property.

 `TemplateName`   
The name of the Elastic Beanstalk configuration template to use with the environment. You must specify either this parameter or a solution stack name.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `Tier`   
Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.

*Required*: No

*Type*: [Elastic Beanstalk Environment Tier Property Type](aws-properties-beanstalk-environment-tier.html "Elastic Beanstalk Environment Tier Property Type")

*Update requires*: See [Elastic Beanstalk Environment Tier Property Type](aws-properties-beanstalk-environment-tier.html "Elastic Beanstalk Environment Tier Property Type")

 `VersionLabel`   
The version to associate with the environment.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `EndpointURL`   
The URL to the load balancer for this environment.

Example:

awseb-myst-myen-132MQC4KRLAMD-1371280482.us-east-1.elb.amazonaws.com

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

### Simple Environment

``` {.programlisting}
        
{
   "Type" : "AWS::ElasticBeanstalk::Environment",
   "Properties" : {
      "ApplicationName" : { "Ref" : "sampleApplication" },
      "Description" :  "AWS Elastic Beanstalk Environment running PHP Sample Application",
      "EnvironmentName" :  "SamplePHPEnvironment",
      "TemplateName" : "DefaultConfiguration",
      "VersionLabel" : "Initial Version"
   }
}        
      
```

### Environment with Embedded Option Settings

``` {.programlisting}
        
{
   "Type" : "AWS::ElasticBeanstalk::Environment",
   "Properties" : {
      "ApplicationName" : { "Ref" : "sampleApplication" },
      "Description" :  "AWS Elastic Beanstalk Environment running Python Sample Application",
      "EnvironmentName" :  "SamplePythonEnvironment",
      "SolutionStackName" : "64bit Amazon Linux running Python",
      "OptionSettings" : [ {
         "Namespace" : "aws:autoscaling:launchconfiguration",
         "OptionName" : "EC2KeyName",
         "Value" : { "Ref" : "KeyName" }
      } ],
      "VersionLabel" : "Initial Version"
   }
}         
      
```

See Also
--------

-   [Launching New Environments](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.environments.html) in the *AWS Elastic Beanstalk Developer Guide*

-   [Managing Environments](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.managing.html) in the *AWS Elastic Beanstalk Developer Guide*

-   For a complete Elastic Beanstalk sample template, see [Elastic Beanstalk Template Snippets](quickref-elasticbeanstalk.html "Elastic Beanstalk Template Snippets").


