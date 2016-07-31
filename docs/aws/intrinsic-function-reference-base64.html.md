Fn::Base64
==========

The intrinsic function `Fn::Base64` returns the Base64 representation of the input string. This function is typically used to pass encoded data to Amazon EC2 instances by way of the UserData property.

Declaration
-----------

{ "Fn::Base64" : *`valueToEncode`* }

Parameters
----------

 valueToEncode   
The string value you want to convert to Base64.

Return Value:
-------------

The original string, in Base64 representation.

Example
-------

``` {.programlisting}
      
         { "Fn::Base64" : "AWS CloudFormation" }
    
```

Supported Functions
-------------------

You can use the `Fn::If` function in the `Fn::Base64` function.

See Also
--------

-   [Intrinsic Function Reference](intrinsic-function-reference.html "Intrinsic Function Reference")


