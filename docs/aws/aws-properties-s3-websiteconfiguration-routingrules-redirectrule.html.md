Amazon S3 Website Configuration Routing Rules Redirect Rule Property
====================================================================

The `RedirectRule` property is an embedded property of the [Amazon S3 Website Configuration Routing Rules Property](aws-properties-s3-websiteconfiguration-routingrules.html "Amazon S3 Website Configuration Routing Rules Property") that describes how requests are redirected. In the event of an error, you can specify a different error code to return.

Syntax
------

``` {.programlisting}
      "RedirectRule" : {
   "HostName" : String,
   "HttpRedirectCode" : String,
   "Protocol" : String,
   "ReplaceKeyPrefixWith" : String,
   "ReplaceKeyWith" : String
}
    
```

Properties
----------

 `                                HostName                            `   
Name of the host where requests are redirected.

*Required*: No

*Type*: String

 `                                HttpRedirectCode                            `   
The HTTP redirect code to use on the response.

*Required*: No

*Type*: String

 `                                Protocol                            `   
The protocol to use in the redirect request.

*Required*: No

*Type*: String

 `                                ReplaceKeyPrefixWith                            `   
The object key prefix to use in the redirect request. For example, to redirect requests for all pages with the prefix `docs/` (objects in the `docs/` folder) to the `documents/` prefix, you can set the `KeyPrefixEquals` property in routing condition property to `docs/`, and set the ReplaceKeyPrefixWith property to `documents/`.

Important

If you specify this property, you cannot specify the `ReplaceKeyWith` property.

*Required*: No

*Type*: String

 `                                ReplaceKeyWith                            `   
The specific object key to use in the redirect request. For example, redirect request to `error.html`.

Important

If you specify this property, you cannot specify the `ReplaceKeyPrefixWith` property.

*Required*: No

*Type*: String


