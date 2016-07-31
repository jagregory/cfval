AWS::ElasticBeanstalk::ApplicationVersion
=========================================

Creates an application version, an iteration of deployable code, for an Elastic Beanstalk application.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::ElasticBeanstalk::ApplicationVersion",
  "Properties" : {
    "ApplicationName" : String,
    "Description" : String,
    "SourceBundle" : { SourceBundle }
  }
}
    
```

Members
-------

 `ApplicationName`   
Name of the Elastic Beanstalk application that is associated with this application version.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Description`   
A description of this application version.

*Required*: No

*Type*: String

*Update requires*: [Some interruptions](using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt)

 `SourceBundle`   
The location of the source bundle for this version.

*Required*: Yes

*Type*: [Source Bundle](aws-properties-beanstalk-sourcebundle.html "Elastic Beanstalk SourceBundle Property Type")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      "myAppVersion" :{ 
  "Type" : "AWS::ElasticBeanstalk::ApplicationVersion",
  "Properties" : {
    "ApplicationName" : {"Ref" : "myApp"},
    "Description" : "my sample version",
    "SourceBundle" : {
      "S3Bucket" : { "Fn::Join" :
        ["-", [ "elasticbeanstalk-samples", { "Ref" : "AWS::Region" } ] ] },
      "S3Key" : "php-sample.zip"
    } 
  }
}
    
```

See Also
--------

-   For a complete Elastic Beanstalk sample template, see [Elastic Beanstalk Template Snippets](quickref-elasticbeanstalk.html "Elastic Beanstalk Template Snippets").


