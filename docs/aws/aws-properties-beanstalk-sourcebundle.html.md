Elastic Beanstalk SourceBundle Property Type
============================================

The SourceBundle property is an embedded property of the [AWS::ElasticBeanstalk::ApplicationVersion](aws-properties-beanstalk-version.html "AWS::ElasticBeanstalk::ApplicationVersion") resource.

Syntax
------

``` {.programlisting}
      
{
   "S3Bucket" : String,
   "S3Key" : String
}     
    
```

Members
-------

 `S3Bucket`   
The Amazon S3 bucket where the data is located.

*Required*: Yes

*Type*: String

 `S3Key`   
The Amazon S3 key where the data is located.

*Required*: Yes

*Type*: String

Example
-------

``` {.programlisting}
      
{
   "S3Bucket" : { "Fn::Join" :
      ["-", [ "elasticbeanstalk-samples", { "Ref" : "AWS::Region" } ] ] },
   "S3Key" : "samplefolder/php-sample.zip"
}     
    
```
