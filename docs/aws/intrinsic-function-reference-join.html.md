Fn::Join
========

The intrinsic function `Fn::Join` appends a set of values into a single value, separated by the specified delimiter. If a delimiter is the empty string, the set of values are concatenated with no delimiter.

Declaration
-----------

"Fn::Join" : [ "*`delimiter`*", [ *`comma-delimited list of          values`* ] ]

Parameters
----------

 delimiter   
The value you want to occur between fragments. The delimiter will occur between fragments only. It will not terminate the final value.

 ListOfValues   
The list of values you want combined.

Return Value
------------

The combined string.

Example
-------

``` {.programlisting}
      "Fn::Join" : [ ":", [ "a", "b", "c" ] ]
    
```

This example returns: "a:b:c".

Supported Functions
-------------------

For the `Fn::Join` delimiter, you cannot use any functions. You must specify a string value.

For the `Fn::Join` list of values, you can use the following functions:

-   `Fn::Base64`

-   `Fn::FindInMap`

-   `Fn::GetAtt`

-   `Fn::GetAZs`

-   `Fn::If`

-   `Fn::Join`

-   `Fn::Select`

-   `Ref`


