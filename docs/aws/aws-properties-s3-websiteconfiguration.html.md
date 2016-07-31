Amazon S3 Website Configuration Property
========================================

WebsiteConfiguration is an embedded property of the [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource.

Syntax
------

``` {.programlisting}
      
"WebsiteConfiguration" : {
   "ErrorDocument" : String,
   "IndexDocument" : String,
   "RedirectAllRequestsTo" : Redirect all requests rule,
   "RoutingRules" : [ Routing rule, ... ]
}     
    
```

Properties
----------

 `ErrorDocument`   
The name of the error document for the website.

*Required*: No

*Type*: String

 `IndexDocument`   
The name of the index document for the website.

*Required*: Yes

*Type*: String

 `RedirectAllRequestsTo`   
The redirect behavior for every request to this bucket's website endpoint.

Important

If you specify this property, you cannot specify any other property.

*Required*: No

*Type*: [Amazon S3 Website Configuration Redirect All Requests To Property](aws-properties-s3-websiteconfiguration-redirectallrequeststo.html "Amazon S3 Website Configuration Redirect All Requests To Property")

 `RoutingRules`   
Rules that define when a redirect is applied and the redirect behavior.

*Required*: No

*Type*: List of [Amazon S3 Website Configuration Routing Rules Property](aws-properties-s3-websiteconfiguration-routingrules.html "Amazon S3 Website Configuration Routing Rules Property")

Example
-------

``` {.programlisting}
      
"S3Bucket" : {
   "Type" : "AWS::S3::Bucket",
   "Properties" : {
      "AccessControl" : "PublicRead",
      "WebsiteConfiguration" : {
         "IndexDocument" : "index.html",
         "ErrorDocument" : "error.html"
      }
   }
}     
    
```

See Also
--------

-   [Custom Error Document Support](http://docs.aws.amazon.com/AmazonS3/latest/dev/CustomErrorDocSupport.html) in the *Amazon Simple Storage Service Developer Guide*

-   [Index Document Support](http://docs.aws.amazon.com/AmazonS3/latest/dev/IndexDocumentSupport.html) in the *Amazon Simple Storage Service Developer Guide*


