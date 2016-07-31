Amazon S3 Cors Configuration Rule
=================================

Describes cross-origin access rules for the [Amazon S3 Cors Configuration](aws-properties-s3-bucket-cors.html "Amazon S3 Cors Configuration") property.

Syntax
------

``` {.programlisting}
      {
  "AllowedHeaders" : [ String, ... ],
  "AllowedMethods" : [ String, ... ],
  "AllowedOrigins" : [ String, ... ],
  "ExposedHeaders" : [ String, ... ],
  "Id" : String,
  "MaxAge" : Integer
}
    
```

Properties
----------

 `AllowedHeaders`   
Headers that are specified in the `Access-Control-Request-Headers` header. These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.

*Required*: No

*Type*: List of strings

 `AllowedMethods`   
An HTTP method that you allow the origin to execute. The valid values are `GET`, `PUT`, `HEAD`, `POST`, and `DELETE`.

*Required*: Yes

*Type*: List of strings

 `AllowedOrigins`   
An origin that you allow to send cross-domain requests.

*Required*: Yes

*Type*: List of strings

 `ExposedHeaders`   
One or more headers in the response that are accessible to client applications (for example, from a JavaScript XMLHttpRequest object).

*Required*: No

*Type*: List of strings

 `Id`   
A unique identifier for this rule. The value cannot be more than 255 characters.

*Required*: No

*Type*: String

 `MaxAge`   
The time in seconds that your browser is to cache the preflight response for the specified resource.

*Required*: No

*Type*: Integer


